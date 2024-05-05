package member

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetClubMembership(db *gorm.DB, userID uuid.UUID) ([]models.Club, *errors.Error) {
	var clubs []models.Club

	user, err := users.GetUser(db, userID)
	if err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("Member").Find(&clubs); err != nil {
		return nil, &errors.FailedToGetUserMemberships
	}

	return clubs, nil
}
