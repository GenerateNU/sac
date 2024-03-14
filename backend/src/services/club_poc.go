package services

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/aws"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ClubPointOfContactServiceInterface interface {
	GetClubPointOfContacts(clubID string) ([]models.PointOfContact, *errors.Error)
	GetClubPointOfContact(clubID, pocID string) (*models.PointOfContact, *errors.Error)
	CreateClubPointOfContact(clubID string, pointOfContactBody models.CreatePointOfContactBody, fileHeader *multipart.FileHeader) (*models.PointOfContact, *errors.Error)
	DeleteClubPointOfContact(clubID, pocID string) *errors.Error
}

type ClubPointOfContactService struct {
	DB       *gorm.DB
	Validate *validator.Validate
	AWSClient *aws.AWSClient
}

func NewClubPointOfContactService(db *gorm.DB, validate *validator.Validate, client *aws.AWSClient) ClubPointOfContactServiceInterface {
	return &ClubPointOfContactService{DB: db, Validate: validate, AWSClient: client}
}

func (cpoc *ClubPointOfContactService) GetClubPointOfContacts(clubID string) ([]models.PointOfContact, *errors.Error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	return transactions.GetClubPointOfContacts(cpoc.DB, *clubIdAsUUID)
}

func (cpoc *ClubPointOfContactService) CreateClubPointOfContact(clubID string, pointOfContactBody models.CreatePointOfContactBody, fileHeader *multipart.FileHeader) (*models.PointOfContact, *errors.Error) {
	if err := cpoc.Validate.Struct(pointOfContactBody); err != nil {
		return nil, &errors.FailedToValidatePointOfContact
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	fileService := NewFileService(cpoc.DB, cpoc.Validate, cpoc.AWSClient)
	file, err := fileService.CreateFile(&models.CreateFileRequestBody{
		OwnerType: "club_point_of_contact",
		OwnerID:   *clubIdAsUUID,
	}, fileHeader)	
	if err != nil {
		return nil, err
	}

	return transactions.CreateClubPointOfContact(cpoc.DB, *clubIdAsUUID, pointOfContactBody, file.ID)
}

func (cpoc *ClubPointOfContactService) GetClubPointOfContact(clubID, pocID string) (*models.PointOfContact, *errors.Error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, &errors.FailedToValidatePointOfContactId
	}

	return transactions.GetClubPointOfContact(cpoc.DB, *pocIdAsUUID, *clubIdAsUUID)
}

func (cpoc *ClubPointOfContactService) DeleteClubPointOfContact(clubID, pocID string) *errors.Error {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return &errors.FailedToValidateClub
	}
	
	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return &errors.FailedToValidatePointOfContactId
	}

	return transactions.DeleteClubPointOfContact(cpoc.DB, *pocIdAsUUID, *clubIdAsUUID)
}