package followers

import (
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubFollowers(db *gorm.DB, clubID uuid.UUID, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Model(&club).Association("Follower").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
}
