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
		return nil, err
	}

	var users []models.User

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Model(&club).Association("Follower").Find(&users); err != nil {
		return nil, &errors.FailedToGetClubFollowers
	}

	return users, nil
}
