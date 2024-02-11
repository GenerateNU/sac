package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubFollowers(db *gorm.DB, clubID uuid.UUID, limit int, page int) ([]models.User, *errors.Error) {
	var users []models.User

	return users, nil
}
