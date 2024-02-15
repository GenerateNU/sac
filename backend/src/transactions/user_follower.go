package transactions

import (
	"slices"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	user, err := GetUser(db, userId)
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubId)
	if err != nil {
		return err
	}

	user.Follower = append(user.Follower, *club)

	if err := db.Model(&user).Association("Follower").Replace(user.Follower); err != nil {
		return &errors.FailedToUpdateUser
	}

	return nil
}

func DeleteFollowing(db *gorm.DB, userId uuid.UUID, clubId uuid.UUID) *errors.Error {
	user, err := GetUser(db, userId, PreloadFollwer())
	if err != nil {
		return err
	}

	club, err := GetClub(db, clubId, PreloadFollwer())
	if err != nil {
		return err
	}

	userFollowingClubIDs := make([]uuid.UUID, len(user.Follower))

	for i, club := range user.Follower {
		userFollowingClubIDs[i] = club.ID
	}

	if !slices.Contains(userFollowingClubIDs, club.ID) {
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
