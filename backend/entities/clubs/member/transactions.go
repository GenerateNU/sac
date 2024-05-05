package member

import (
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/entities/users/follower"
	"github.com/GenerateNU/sac/backend/errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubMembers(db *gorm.DB, clubID uuid.UUID, limit int, page int) ([]models.User, *errors.Error) {
	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return nil, err
	}

	var users []models.User

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Model(&club).Association("Member").Find(&users); err != nil {
		return nil, &errors.FailedToGetClubMembers
	}

	return users, nil
}

func CreateClubMember(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) *errors.Error {
	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := clubs.GetClub(db, clubID)
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

	if err := follower.CreateFollowing(tx, userID, clubID); err != nil {
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

func DeleteClubMember(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) *errors.Error {
	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Model(&user).Association("Member").Delete(club); err != nil {
		tx.Rollback()
		return &errors.FailedToUpdateUser
	}

	if err := follower.DeleteFollowing(tx, userID, clubID); err != nil {
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
