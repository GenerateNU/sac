package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateMember(db *gorm.DB, userID uuid.UUID, clubID uuid.UUID) *errors.Error {
	user, err := GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubID)
	if err != nil {
		return err
	}

	tx := db.Begin()

	var count int64
	if err := tx.Model(&models.Membership{}).Where("user_id = ? AND club_id = ?", userID, clubID).Count(&count).Error; err != nil {
		return &errors.FailedToGetUserMemberships
	}

	if count > 0 {
		return nil
	}

	if err := tx.Model(&user).Association("Member").Append(club); err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
	}

	if err := CreateFollowing(tx, userID, clubID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&club).Update("num_members", gorm.Expr("num_members + 1")).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
	}

	if err := tx.Commit().Error; err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func DeleteMember(db *gorm.DB, userID uuid.UUID, clubID uuid.UUID) *errors.Error {
	user, err := GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubID)
	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Model(&user).Association("Member").Delete(club); err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
	}

	if err := DeleteFollowing(tx, userID, clubID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&club).Update("num_members", gorm.Expr("num_members - 1")).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
	}

	if err := tx.Commit().Error; err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func GetClubMembership(db *gorm.DB, userID uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	user, err := GetUser(db, userID)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Member").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserMemberships
	}

	return clubs, nil
}
