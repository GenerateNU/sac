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
	if err := db.WithContext(ctx).Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Append(recruitment); err != nil {
		return nil, err
	}

	return recruitment, nil
}

func GetClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID) (*models.Recruitment, error) {
	var recruitment models.Recruitment
	if err := db.WithContext(ctx).Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Find(&recruitment); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}

		return nil, err
	}

	return &recruitment, nil
}

func UpdateClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Recruitment, error) {
	existingRecruitment, err := GetClubRecruitment(ctx, db, clubID)
	if err != nil {
		return nil, err
	}

	if err := db.WithContext(ctx).Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Replace(utilities.Merge(existingRecruitment, recruitment)); err != nil {
		return nil, err
	}

	return recruitment, nil
}

func DeleteClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID) error {
	if err := db.WithContext(ctx).Model(&models.Club{}).Where("id = ?", clubID).Association("Recruitment").Clear(); err != nil {
		return err
	}

	return nil
}

func CreateApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, application *models.Application) (*models.Application, error) {
	recruitment, err := GetClubRecruitment(ctx, db, clubID)
	if err != nil {
		return nil, err
	}

	if err := db.WithContext(ctx).Model(&recruitment).Association("Application").Append(application); err != nil {
		return nil, err
	}

	return application, nil
}

func GetClubRecruitmentApplications(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, pageInfo fiberpaginate.PageInfo) ([]models.Application, error) {
	recruitment, err := GetClubRecruitment(ctx, db, clubID)
	if err != nil {
		return nil, err
	}

	var applications []models.Application
	if err := db.WithContext(ctx).Scopes(utilities.IntoScope(pageInfo, db)).Model(&recruitment).Association("Application").Find(&applications); err != nil {
		return nil, err
	}

	return applications, nil
}

func GetClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, applicationID *uuid.UUID) (*models.Application, error) {
	var application models.Application
	if err := db.WithContext(ctx).Model(&models.Recruitment{}).Where("id = ?", recruitmentID).Association("Application").Find(&application, applicationID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}

		return nil, err
	}

	return &application, nil
}

func UpdateClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, applicationID *uuid.UUID, application *models.Application) (*models.Application, error) {
	existingApplication, err := GetClubRecruitmentApplication(ctx, db, clubID, recruitmentID, applicationID)
	if err != nil {
		return nil, err
	}

	if err := db.WithContext(ctx).Model(&models.Recruitment{}).Where("id = ?", recruitmentID).Association("Application").Replace(utilities.Merge(existingApplication, application)); err != nil {
		return nil, err
	}

	return application, nil
}

func DeleteClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitmentID *uuid.UUID, applicationID *uuid.UUID) error {
	if err := db.WithContext(ctx).Model(&models.Recruitment{}).Where("id = ?", recruitmentID).Association("Application").Delete(applicationID); err != nil {
		return err
	}

	return nil
}
