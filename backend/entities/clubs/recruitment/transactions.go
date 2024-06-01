package recruitment

import (
	"context"
	"errors"

	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Recruitment, error) {
	club, err := clubs.GetClub(db, *clubID, transactions.PreloadRecruitment())
	if err != nil {
		return nil, err
	}

	if err := db.WithContext(ctx).Model(&club).Association("Recruitment").Replace(recruitment); err != nil {
		return nil, err
	}

	return recruitment, nil
}

func GetClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID) (*models.Recruitment, error) {
	club, err := clubs.GetClub(db, *clubID)
	if err != nil {
		return nil, err
	}

	var recruitment models.Recruitment
	if err := db.WithContext(ctx).Model(&club).Association("Recruitment").Find(&recruitment); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}

		return nil, err
	}

	return &recruitment, nil
}

func UpdateClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, recruitment *models.Recruitment) (*models.Recruitment, error) {
	club, err := clubs.GetClub(db, *clubID, transactions.PreloadRecruitment())
	if err != nil {
		return nil, err
	}

	if err := db.WithContext(ctx).Model(&club).Association("Recruitment").Replace(utilities.Merge(&club.Recruitment, recruitment)); err != nil {
		return nil, err
	}

	return recruitment, nil
}

func DeleteClubRecruitment(ctx context.Context, db *gorm.DB, clubID *uuid.UUID) error {
	club, err := clubs.GetClub(db, *clubID, transactions.PreloadRecruitment())
	if err != nil {
		return err
	}

	if err := db.WithContext(ctx).Model(&club).Association("Recruitment").Clear(); err != nil {
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

func GetClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, applicationID *uuid.UUID) (*models.Application, error) {
	club, err := clubs.GetClub(db, *clubID, transactions.PreloadRecruitment())
	if err != nil {
		return nil, err
	}

	var application models.Application
	if err := db.WithContext(ctx).Model(&club.Recruitment).Association("Application").Find(&application, applicationID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}

		return nil, err
	}

	return &application, nil
}

func UpdateClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, applicationID *uuid.UUID, application *models.Application) (*models.Application, error) {
	existingApplication, err := GetClubRecruitmentApplication(ctx, db, clubID, applicationID)
	if err != nil {
		return nil, err
	}

	if err := db.WithContext(ctx).Model(existingApplication).Updates(application).Error; err != nil {
		return nil, err
	}

	return application, nil
}

func DeleteClubRecruitmentApplication(ctx context.Context, db *gorm.DB, clubID *uuid.UUID, applicationID *uuid.UUID) error {
	application, err := GetClubRecruitmentApplication(ctx, db, clubID, applicationID)
	if err != nil {
		return err
	}

	if err := db.WithContext(ctx).Delete(application).Error; err != nil {
		return err
	}

	return nil
}
