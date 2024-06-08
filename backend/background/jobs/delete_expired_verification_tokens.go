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

func (j *Jobs) DeleteExpiredVerificationTokens(ctx context.Context) background.JobFunc {
	return func() {
		t := time.NewTicker(constants.DELETE_EXPIRED_VERIFICATION_INTERVAL)

		for range t.C {
			slog.Info("deleting expired verification tokens")
			if err := deleteExpiredVerificationTokens(ctx, j.db, constants.DELETE_EXPIRED_TOKEN_LIMIT); err != nil {
				slog.Error("failed to delete expired verification tokens", "err", err)
			}
		}
	}
}

func deleteExpiredVerificationTokens(ctx context.Context, db *gorm.DB, limit int) error {
	result := db.WithContext(ctx).Model(&models.Verification{}).Where("expires_at < ?", time.Now().UTC()).Limit(limit).Delete(&models.Verification{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		slog.Info("no expired verification tokens found")
	}

	return nil
}
