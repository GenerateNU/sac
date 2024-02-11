package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	user, err := GetUser(db, userId)
	if err != nil {
		return &errors.UserNotFound
	}

	club, err := GetClub(db, clubId)
	if err != nil {
		return &errors.ClubNotFound
	}

	user.Follower = append(user.Follower, *club)

	if err := db.Model(&user).Association("Follower").Append(&club); err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func DeleteFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	user, err := GetUser(db, userId)
	if err != nil {
		return &errors.UserNotFound
	}
	club, err := GetClub(db, clubId)
	if err != nil {
		return &errors.ClubNotFound
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
		return nil, &errors.UserNotFound
	}

	if err := db.Model(&user).Association("Follower").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserFollowing
	}
	return clubs, nil
}
