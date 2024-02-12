package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ContactServiceInterface interface {
	GetContacts(limit string, page string) ([]models.Contact, *errors.Error)
	GetContact(contactID string) (*models.Contact, *errors.Error)
	DeleteContact(contactID string) *errors.Error
}

type ContactService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewContactService(db *gorm.DB, validate *validator.Validate) *ContactService {
	return &ContactService{DB: db, Validate: validate}
}

func (c *ContactService) GetContacts(limit string, page string) ([]models.Contact, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetContacts(c.DB, *limitAsInt, offset)
}

func (c *ContactService) GetContact(contactID string) (*models.Contact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(contactID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetContact(c.DB, *idAsUUID)
}

func (c *ContactService) DeleteContact(contactID string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(contactID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteContact(c.DB, *idAsUUID)
}
