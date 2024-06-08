package jobs

import (
	"context"
	"log/slog"
	"time"

	"github.com/GenerateNU/sac/backend/background"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"gorm.io/gorm"
)

func (j *Jobs) VerificationSender(ctx context.Context) background.JobFunc {
	return func() {
		t := time.NewTicker(constants.EMAIL_SENDER_INTERVAL)

		for range t.C {
			slog.Info("sending verification email")

			func() {
				// start a separate transaction to increment the attempts
				incrementTx := j.db.WithContext(ctx).Begin()
				defer func() {
					if r := recover(); r != nil {
						incrementTx.Rollback()
					}
				}()

				task, err := j.dequeueVerificationTask(incrementTx)
				if err != nil {
					incrementTx.Rollback()
					return
				}

				// increment attempts
				task.Attempts++
				if err := incrementTx.Save(task).Error; err != nil {
					slog.Error("failed to increment attempts", "err", err)
					incrementTx.Rollback()
					return
				}

				if err := incrementTx.Commit().Error; err != nil {
					slog.Error("failed to commit increment transaction", "err", err)
					return
				}

				// start the main transaction
				tx := j.db.WithContext(ctx).Begin()
				defer func() {
					if r := recover(); r != nil {
						tx.Rollback()
					}
				}()

				// attempt to send the email
				err = j.sendVerificationEmail(ctx, tx, task)
				if err != nil {
					slog.Error("failed to send verification email", "err", err)
					if task.Attempts >= constants.MAX_EMAIL_ATTEMPTS {
						if err := j.deleteVerificationTask(tx, task); err != nil {
							slog.Error("failed to delete verification task", "err", err)
						}
					}
				} else {
					if err := j.deleteVerificationTask(tx, task); err != nil {
						slog.Error("failed to delete verification task", "err", err)
					}
				}

				if err := tx.Commit().Error; err != nil {
					slog.Error("failed to commit transaction", "err", err)
				}
			}()
		}
	}
}

func (j *Jobs) dequeueVerificationTask(tx *gorm.DB) (*models.VerificationTask, error) {
	var task models.VerificationTask
	if err := tx.Raw("SELECT user_id, verification_id, attempts FROM verification_tasks FOR UPDATE SKIP LOCKED LIMIT 1").Scan(&task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &task, nil
}

func (j *Jobs) deleteVerificationTask(tx *gorm.DB, task *models.VerificationTask) error {
	return tx.Delete(task).Error
}

func (j *Jobs) sendVerificationEmail(ctx context.Context, tx *gorm.DB, task *models.VerificationTask) error {
	var verification models.Verification
	if err := tx.Model(&models.Verification{}).Where("id = ?", task.VerificationID).First(&verification).Error; err != nil {
		return err
	}

	return j.emailer.SendVerification(ctx, verification.Email, verification.Token)
}
