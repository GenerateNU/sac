package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"gorm.io/gorm"
)

type ClubFollowerServiceInterface interface {
	GetClubFollowers(clubID string, limit string, page string) ([]models.User, *errors.Error)
}

type ClubFollowerService struct {
	DB *gorm.DB
}

func NewClubFollowerService(db *gorm.DB) *ClubFollowerService {
	return &ClubFollowerService{DB: db}
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
