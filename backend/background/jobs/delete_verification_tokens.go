package jobs

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/GenerateNU/sac/backend/background"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (j *Jobs) DeleteVerificationTokens(ctx context.Context) background.JobFunc {
	return func() {
		t := time.NewTicker(constants.DELETE_VERIFICATION_INTERVAL)

		for range t.C {
			slog.Info("deleting verification token")
			func() {
				tx := j.db.WithContext(ctx).Begin()
				defer func() {
					if r := recover(); r != nil {
						tx.Rollback()
					}
				}()

				task, err := dequeTask(ctx, tx)
				if err != nil {
					slog.Error("failed to fetch task", "err", err)
					tx.Rollback()
				}

				if task == nil {
					tx.Rollback()
					return
				}

				if err := deleteTask(ctx, tx, *task); err != nil {
					slog.Error("failed to delete task", "err", err)
					tx.Rollback()
				}
			}()
		}
	}
}

type task struct {
	token  string    `gorm:"column:verification_token"`
	userID uuid.UUID `gorm:"column:user_id"`
}

func dequeTask(ctx context.Context, tx *gorm.DB) (*task, error) {
	var task task
	if err := tx.WithContext(ctx).Raw(
		`
	SELECT verification_token, user_id
	FROM verifications_deletion_queue
	FOR UPDATE
	SKIP LOCKED
	LIMIT 1
	`,
	).Scan(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &task, nil
}

func deleteTask(ctx context.Context, tx *gorm.DB, task task) error {
	if err := tx.WithContext(ctx).Where("verification_token = ? AND user_id = ?", task.token, task.userID).Delete(&models.VerificationDeletionQueue{}).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}
