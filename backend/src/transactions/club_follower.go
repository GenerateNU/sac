package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubFollowers(db *gorm.DB, clubID uuid.UUID, limit int, page int) ([]models.User, *errors.Error) {
	club, err := GetClub(db, clubID)
	if err != nil {
		return nil, &errors.ClubNotFound
	}

	var users []models.User

	if err := db.Model(&club).Association("Followers").Find(&users); err != nil {
		return nil, &errors.FailedToGetClubFollowers
	}

	return users, nil
}
