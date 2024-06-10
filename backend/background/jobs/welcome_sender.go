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

func (j *Jobs) WelcomeSender(ctx context.Context) background.JobFunc {
	return func() {
		t := time.NewTicker(constants.EMAIL_SENDER_INTERVAL)

		for range t.C {
			slog.Info("sending welcome email")

			func() {
				// start a separate transaction to increment the attempts
				incrementTx := j.db.WithContext(ctx).Begin()
				defer func() {
					if r := recover(); r != nil {
						incrementTx.Rollback()
					}
				}()

				task, err := j.dequeueWelcomeTask(incrementTx)
				if err != nil {
					incrementTx.Rollback()
					return
				}

				if task == nil {
					slog.Info("no welcome tasks to send")
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
				err = j.sendWelcomeEmail(ctx, task)
				if err != nil {
					slog.Error("failed to send welcome email", "err", err)
					if task.Attempts >= constants.MAX_EMAIL_ATTEMPTS {
						if err := j.deleteWelcomeTask(tx, task); err != nil {
							slog.Error("failed to delete welcome task", "err", err)
						}
					}
				} else {
					if err := j.deleteWelcomeTask(tx, task); err != nil {
						slog.Error("failed to delete welcome task", "err", err)
					}
				}

				if err := tx.Commit().Error; err != nil {
					slog.Error("failed to commit transaction", "err", err)
				}
			}()
		}
	}
}

func (j *Jobs) dequeueWelcomeTask(tx *gorm.DB) (*models.WelcomeTask, error) {
	var task models.WelcomeTask
	if err := tx.Raw("SELECT email, name, attempts FROM welcome_tasks FOR UPDATE SKIP LOCKED LIMIT 1").Scan(&task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if task.Email == "" && task.Name == "" {
		tx.Rollback()
		return nil, nil
	}

	return &task, nil
}

func (j *Jobs) deleteWelcomeTask(tx *gorm.DB, task *models.WelcomeTask) error {
	return tx.Delete(task).Error
}

func (j *Jobs) sendWelcomeEmail(ctx context.Context, task *models.WelcomeTask) error {
	return j.emailer.SendWelcome(ctx, task.Email, task.Name)
}
