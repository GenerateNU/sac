package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
)

type ClubServiceInterface interface {
	GetClubs(limit string, page string) ([]models.Club, *errors.Error)
	GetClub(id string) (*models.Club, *errors.Error)
	CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, *errors.Error)
	UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, *errors.Error)
	DeleteClub(id string) *errors.Error
	GetContacts(limit string, page string) ([]models.Contact, *errors.Error)
	GetClubContacts(id string) ([]models.Contact, *errors.Error)
	PutContact(clubID string, contactBody models.PutContactRequestBody) (*models.Contact, *errors.Error)
	DeleteContact(id string) *errors.Error
}

type ClubService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (c *ClubService) GetClubs(limit string, page string) ([]models.Club, *errors.Error) {
	limitAsInt, err := utilities.ValidateNonNegative(limit)

	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)

	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	offset := (*pageAsInt - 1) * *limitAsInt

	return transactions.GetClubs(c.DB, *limitAsInt, offset)
}

func (c *ClubService) CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, *errors.Error) {
	if err := c.Validate.Struct(clubBody); err != nil {
		return nil, &errors.FailedToValidateClub
	}

	club, err := utilities.MapRequestToModel(clubBody, &models.Club{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.CreateClub(c.DB, clubBody.UserID, *club)
}

func (c *ClubService) GetClub(id string) (*models.Club, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClub(c.DB, *idAsUUID)
}

func (c *ClubService) UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(id)
	if idErr != nil {
		return nil, idErr
	}

	if err := c.Validate.Struct(clubBody); err != nil {
		return nil, &errors.FailedToValidateClub
	}

	club, err := utilities.MapRequestToModel(clubBody, &models.Club{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}

	return transactions.UpdateClub(c.DB, *idAsUUID, *club)
}

func (c *ClubService) DeleteClub(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

return transactions.DeleteClub(c.DB, *idAsUUID)
}

func (c *ClubService) GetContacts(limit string, page string) ([]models.Contact, *errors.Error) {
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

func (c *ClubService) GetClubContacts(id string) ([]models.Contact, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClubContacts(c.DB, *idAsUUID)
}

func (c *ClubService) PutContact(clubID string, contactBody models.PutContactRequestBody) (*models.Contact, *errors.Error) {
	idAsUUID, idErr := utilities.ValidateID(clubID)
	if idErr != nil {
		return nil, idErr
	}

	if err := c.Validate.Struct(contactBody); err != nil {
		return nil, &errors.FailedToValidateContact
	}

	contact, err := utilities.MapRequestToModel(contactBody, &models.Contact{})
	if err != nil {
		return nil, &errors.FailedToMapRequestToModel
	}
	contact.ClubID = *idAsUUID
	return transactions.PutContact(c.DB, *idAsUUID, *contact)
}

func (c *ClubService) DeleteContact(id string) *errors.Error {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteContact(c.DB, *idAsUUID)
}
