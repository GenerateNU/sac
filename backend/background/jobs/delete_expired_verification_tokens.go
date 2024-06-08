package jobs

import (
	"context"
	"log/slog"
	"time"

	"github.com/GenerateNU/sac/backend/background"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/entities/models"
)

func (j *Jobs) ExpiredVerificationTokensDeleter(ctx context.Context) background.JobFunc {
	return func() {
		t := time.NewTicker(constants.DELETE_EXPIRED_VERIFICATION_INTERVAL)

		for range t.C {
			slog.Info("deleting expired verification tokens")
			if err := j.deleteExpiredTokens(ctx, constants.DELETE_EXPIRED_VERIFICATION_LIMIT); err != nil {
				slog.Error("failed to delete expired verification tokens", "err", err)
			}
		}
	}
}

func (j *Jobs) deleteExpiredTokens(ctx context.Context, limit int) error {
	result := j.db.WithContext(ctx).Model(&models.Verification{}).Where("expires_at < ?", time.Now().UTC()).Limit(limit).Delete(&models.Verification{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		slog.Info("no expired verification tokens found")
	}

	return nil
}
