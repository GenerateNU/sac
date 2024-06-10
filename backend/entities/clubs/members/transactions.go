package members

import (
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/clubs/followers"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubMembers(db *gorm.DB, clubID uuid.UUID, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Model(&club).Association("Member").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func CreateClubMember(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) error {
	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var count int64
	if err := tx.Model(&models.Membership{}).Where("user_id = ? AND club_id = ?", userID, clubID).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	if err := tx.Model(&user).Association("Member").Append(club); err != nil {
		tx.Rollback()
		return err
	}

	if err := followers.CreateClubFollower(tx, userID, clubID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&club).Update("num_members", gorm.Expr("num_members + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteClubMember(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) error {
	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&user).Association("Member").Delete(club); err != nil {
		tx.Rollback()
		return err
	}

	if err := followers.DeleteClubFollower(tx, userID, clubID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&club).Update("num_members", gorm.Expr("num_members - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
