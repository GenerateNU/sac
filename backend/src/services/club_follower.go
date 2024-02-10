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
	GetClubFollowers(clubID string, limit string, page string) ([]models.User, *errors.Error)
}

type ClubFollowerService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewClubFollowerService(db *gorm.DB, validate *validator.Validate) *ClubFollowerService {
	return &ClubFollowerService{DB: db, Validate: validate}
}

func (cf *ClubFollowerService) GetClubFollowers(clubID string, limit string, page string) ([]models.User, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetClubFollowers(cf.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
