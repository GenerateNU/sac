package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubs(db *gorm.DB, pageInfo fiberpaginate.PageInfo) ([]models.Club, error) {
	query := db.Model(&models.Club{})
	var clubs []models.Club

	if err := query.Scopes(utilities.IntoScope(pageInfo, query)).Find(&clubs).Error; err != nil {
		return nil, err
	}

	return clubs, nil
}

func CreateClub(db *gorm.DB, userId uuid.UUID, club models.Club) (*models.Club, error) {
	user, err := users.GetUser(db, userId)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()

	club.NumMembers = 1

	if err := tx.Create(&club).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	membership := models.Membership{
		ClubID:         club.ID,
		UserID:         user.ID,
		MembershipType: models.MembershipTypeAdmin,
	}

	if err := tx.Create(&membership).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&user).Association("Follower").Append(&club); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &club, nil
}

func UpdateClub(db *gorm.DB, id uuid.UUID, club models.Club) (*models.Club, error) {
	tx := db.Begin()

	var existingClub models.Club

	err := tx.First(&existingClub, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	if err := tx.Model(&existingClub).Updates(&club).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &existingClub, nil
}

func DeleteClub(db *gorm.DB, id uuid.UUID) error {
	tx := db.Begin()

	var existingClub models.Club
	err := tx.First(&existingClub, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utilities.ErrNotFound
		}
		return err
	}

	if result := tx.Delete(&models.Club{}, id); result.RowsAffected == 0 {
		tx.Rollback()
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
