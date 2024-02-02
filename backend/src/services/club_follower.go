package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ClubFollowerServiceInterface interface {
	GetUserFollowingClubs(userID string) ([]models.Club, *errors.Error)
}

type ClubFollowerService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewClubFollowerService(db *gorm.DB, validate *validator.Validate) *ClubFollowerService {
	return &ClubFollowerService{DB: db, Validate: validate}
}

func (cf *ClubFollowerService) GetUserFollowingClubs(userID string) ([]models.Club, *errors.Error) {
	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return nil, err
	}

	return transactions.GetClubFollowing(cf.DB, *userIDAsUUID)
}
