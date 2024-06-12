package followers

import (
	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubFollowers(db *gorm.DB, clubID uuid.UUID, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	db = cache.SetCache(db, true)

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

func CreateClubFollower(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) error {
	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return err
	}

	if err := db.Model(&user).Association("Follower").Append(club); err != nil {
		return err
	}

	return nil
}

func DeleteClubFollower(db *gorm.DB, clubID uuid.UUID, userID uuid.UUID) error {
	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	club, err := clubs.GetClub(db, clubID)
	if err != nil {
		return err
	}

	if err := db.Model(&user).Association("Follower").Delete(club); err != nil {
		return err
	}

	return nil
}
