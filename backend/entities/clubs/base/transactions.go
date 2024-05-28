package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	search "github.com/GenerateNU/sac/backend/search/base"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

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

	err = search.Upsert[models.Club](db, constants.CLUBS_INDEX, club.ID.String(), &club)
	if err != nil {
		return nil, err
	}

	return &club, nil
}

func UpdateClub(db *gorm.DB, id uuid.UUID, club models.Club) (*models.Club, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingClub models.Club
	err := tx.First(&existingClub, id).Error
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	if err := tx.Model(&existingClub).Clauses(clause.Returning{}).Updates(&club).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	err = search.Upsert[models.Club](db, constants.CLUBS_INDEX, club.ID.String(), &club)
	if err != nil {
		return nil, err
	}

	return &existingClub, tx.Commit().Error
}

func DeleteClub(db *gorm.DB, id uuid.UUID) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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

	err = search.Delete(db, constants.CLUBS_INDEX, id.String())
	if err != nil {
		return err
	}

	return tx.Commit().Error
}
