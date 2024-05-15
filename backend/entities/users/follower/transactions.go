package follower

import (
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFollowing(db *gorm.DB, userID uuid.UUID, clubID uuid.UUID) error {
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

func DeleteFollowing(db *gorm.DB, userID uuid.UUID, clubID uuid.UUID) error {
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

func GetClubFollowing(db *gorm.DB, userID uuid.UUID) ([]models.Club, error) {
	var clubs []models.Club

	user, err := users.GetUser(db, userID)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Follower").Find(&clubs); err != nil {
		return nil, err
	}

	return clubs, nil
}
