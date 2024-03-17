package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type PointOfContactServiceInterface interface {
	GetPointOfContacts(limit string, page string) ([]models.PointOfContact, *errors.Error)
	GetPointOfContact(pocID string) (*models.PointOfContact, *errors.Error)
}

type PointOfContactService struct {
	DB *gorm.DB
	Validate *validator.Validate
}

func NewPointOfContactService(db *gorm.DB, validate *validator.Validate) *PointOfContactService {
	return &PointOfContactService{DB: db, Validate: validate}
}

func (poc *PointOfContactService) GetPointOfContacts(limit string, page string) ([]models.PointOfContact, *errors.Error) {	
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetPointOfContacts(poc.DB, *limitAsInt, *pageAsInt)
}

func (poc *PointOfContactService) GetPointOfContact(pocID string) (*models.PointOfContact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetPointOfContact(poc.DB, *idAsUUID)
}
