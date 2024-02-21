package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	tx := db.Begin()

	user, err := GetUser(tx, userId)
	if err != nil {
		return err
	}

	club, err := GetClub(tx, clubId)
	if err != nil {
		return err
	}

	var count int64
	if err := tx.Model(&models.Follower{}).Where("user_id = ? AND club_id = ?", userId, clubId).Count(&count).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToGetUserFollowing
	}

	if count > 0 {
		tx.Rollback()
		return &errors.UserAlreadyFollowingClub
	}

	if err := db.Model(&user).Association("Follower").Append(club); err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func DeleteFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	tx := db.Begin()

	user, err := GetUser(tx, userId)
	if err != nil {
		return err
	}

	club, err := GetClub(tx, clubId)
	if err != nil {
		return err
	}

	var count int64
	if err := tx.Model(&models.Follower{}).Where("user_id = ? AND club_id = ?", userId, clubId).Count(&count).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToGetUserFollowing
	}

	if count == 0 {
		tx.Rollback()
		return &errors.UserNotFollowingClub
	}

	if err := db.Model(&user).Association("Follower").Delete(club); err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func GetClubFollowing(db *gorm.DB, userId uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	user, err := GetUser(db, userId)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Follower").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserFollowing
	}

	return clubs, nil
}
