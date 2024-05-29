package recruitment

import (
	"context"
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Recruitment, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Recruitment, error) {
		if err := db.Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Append(recruitment); err != nil {
			return nil, err
		}

		return recruitment, nil
	})
}

func GetClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID) (*models.Recruitment, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Recruitment, error) {
		var recruitment models.Recruitment
		if err := db.Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Find(&recruitment); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, utilities.ErrNotFound
			}

			return nil, err
		}

		return &recruitment, nil
	})
}

func UpdateClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Recruitment, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Recruitment, error) {
		existingRecruitment, err := GetClubRecruitment(ctx, db, clubID)
		if err != nil {
			return nil, err
		}

		if err := db.Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Replace(utilities.Merge(existingRecruitment, recruitment)); err != nil {
			return nil, err
		}

		return recruitment, nil
	})
}

func DeleteClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID) error {
	return utilities.ExecuteWithTimeout(context.Background(), func() error {
		if err := db.Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Clear(); err != nil {
			return err
		}

		return nil
	})
}

func CreateApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, application *models.Application) (*models.Application, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Application, error) {
		recruitment, err := GetClubRecruitment(ctx, db, clubID)
		if err != nil {
			return nil, err
		}

		if err := db.Model(&recruitment).Association("Application").Append(application); err != nil {
			return nil, err
		}

		return application, nil
	})
}

func GetClubRecruitmentApplications(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, pageInfo fiberpaginate.PageInfo) ([]models.Application, error) {
	result, err := utilities.ExecuteWithTimeoutResult(ctx, func() (*[]models.Application, error) {
		recruitment, err := GetClubRecruitment(ctx, db, clubID)
		if err != nil {
			return nil, err
		}

		var applications []models.Application
		if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Model(&recruitment).Association("Application").Find(&applications); err != nil {
			return nil, err
		}

		return &applications, nil
	})

	return *result, err
}

func GetClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, applicationID *uuid.UUID) (*models.Application, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Application, error) {
		var application models.Application
		if err := db.Model(&models.Recruitment{}).Where("id = ?", recruitmentID).Association("Application").Find(&application, applicationID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, utilities.ErrNotFound
			}

			return nil, err
		}

		return &application, nil
	})
}

func UpdateClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, applicationID *uuid.UUID, application *models.Application) (*models.Application, error) {
	return utilities.ExecuteWithTimeoutResult(ctx, func() (*models.Application, error) {
		existingApplication, err := GetClubRecruitmentApplication(ctx, db, clubID, recruitmentID, applicationID)
		if err != nil {
			return nil, err
		}

		if err := db.Model(&models.Recruitment{}).Where("id = ?", recruitmentID).Association("Application").Replace(utilities.Merge(existingApplication, application)); err != nil {
			return nil, err
		}

		return application, nil
	})
}

func DeleteClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, applicationID *uuid.UUID) error {
	return utilities.ExecuteWithTimeout(ctx, func() error {
		if err := db.Model(&models.Recruitment{}).Where("id = ?", recruitmentID).Association("Application").Delete(applicationID); err != nil {
			return err
		}

		return nil
	})
}
