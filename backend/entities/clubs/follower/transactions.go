package follower

import (
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubFollowers(db *gorm.DB, clubID uuid.UUID, limit int, page int) ([]models.User, error) {
	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return nil, err
	}

	var users []models.User

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Model(&club).Association("Follower").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
}
