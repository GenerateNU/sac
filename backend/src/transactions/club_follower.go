package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserFollowingClubs(db *gorm.DB, club_id uuid.UUID) ([]models.User, *errors.Error) {
	var users []models.User
	club, err := GetClub(db, club_id)
	if err != nil {
		return nil, &errors.ClubNotFound
	}

	if err := db.Model(&club).Association("Follower").Find(&users); err != nil {
		return nil, &errors.FailedToGetClubFollowers
	}
	return users, nil
}
