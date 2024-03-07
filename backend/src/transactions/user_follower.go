package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFollowing(db *gorm.DB, userID uuid.UUID, clubID uuid.UUID) *errors.Error {
	user, err := GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubID)
	if err != nil {
		return err
	}

	if err := db.Model(&user).Association("Follower").Append(club); err != nil {
		return &errors.FailedToFollowClub
	}

	return nil
}

func DeleteFollowing(db *gorm.DB, userID uuid.UUID, clubID uuid.UUID) *errors.Error {
	user, err := GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubID)
	if err != nil {
		return err
	}

	if err := db.Model(&user).Association("Follower").Delete(club); err != nil {
		return &errors.UserNotFollowingClub
	}

	return nil
}

func GetClubFollowing(db *gorm.DB, userID uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	user, err := GetUser(db, userID)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Follower").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserFollowing
	}

	return clubs, nil
}
