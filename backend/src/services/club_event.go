package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"gorm.io/gorm"
)

type ClubEventServiceInterface interface {
	GetClubEvents(clubID string, limit string, page string) ([]models.Event, *errors.Error)
}

type ClubEventService struct {
	DB *gorm.DB
}

func NewClubEventService(db *gorm.DB) *ClubEventService {
	return &ClubEventService{DB: db}
}

func (c *ClubEventService) GetClubEvents(clubID string, limit string, page string) ([]models.Event, *errors.Error) {
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

	return transactions.GetClubEvents(c.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
