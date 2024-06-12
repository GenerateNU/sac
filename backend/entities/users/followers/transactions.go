package followers

import (
	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubFollowing(db *gorm.DB, userID uuid.UUID) ([]models.Club, error) {
	db = cache.SetCache(db, true)

	user, err := users.GetUser(db, userID)
	if err != nil {
		return nil, err
	}

	var clubs []models.Club
	if err := db.Model(&user).Association("Follower").Find(&clubs); err != nil {
		return nil, err
	}

	return clubs, nil
}
