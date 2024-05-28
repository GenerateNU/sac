package recruitment

import (
	"context"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Club, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Club, error) {
		return nil, nil
	})
}

func UpdateClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Club, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Club, error) {
		return nil, nil
	})
}

func DeleteClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID) error {
	return utilities.ExecuteWithTimeout(context.Background(), func() error {
		return nil
	})
}
