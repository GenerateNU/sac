package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateMember(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	tx := db.Begin()

	user, err := GetUser(tx, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	club, err := GetClub(tx, clubId)
	if err != nil {
		tx.Rollback()
		return err
	}

	var count int64
	if err := tx.Model(&models.Membership{}).Where("user_id = ? AND club_id = ?", userId, clubId).Count(&count).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToGetUserMemberships
	}

	if count > 0 {
		tx.Rollback()
		return &errors.AlreadyMemberOfClub
	}

	if err := tx.Model(&user).Association("Member").Append(club); err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
	}

	if err := tx.Model(&user).Association("Follower").Append(club); err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
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

func DeleteMember(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	tx := db.Begin()

	user, err := GetUser(tx, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	club, err := GetClub(tx, clubId)
	if err != nil {
		tx.Rollback()
		return err
	}

	var count int64
	if err := tx.Model(&models.Membership{}).Where("user_id = ? AND club_id = ?", userId, clubId).Count(&count).Error; err != nil {
		tx.Rollback()
		return &errors.FailedToGetUserMemberships
	}

	if count == 0 {
		tx.Rollback()
		return &errors.UserNotMemberOfClub
	}

	if err := tx.Model(&user).Association("Member").Delete(club); err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
	}

	if err := tx.Model(&user).Association("Follower").Delete(club); err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
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

func GetClubMembership(db *gorm.DB, userId uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	user, err := GetUser(db, userId)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Member").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserMemberships
	}

	return clubs, nil
}
