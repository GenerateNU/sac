package services

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/file"
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
	AWSProvider *file.AWSProvider
}

func NewClubPointOfContactService(db *gorm.DB, validate *validator.Validate, client *file.AWSProvider) ClubPointOfContactServiceInterface {
	return &ClubPointOfContactService{DB: db, Validate: validate, AWSProvider: client}
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

	fileInfo, err := cpoc.AWSProvider.UploadFile("point_of_contacts", fileHeader)
	if err != nil {
		return nil, err
	}

	tx := cpoc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	poc, err := transactions.CreateClubPointOfContact(tx, *clubIdAsUUID, pointOfContactBody)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = transactions.CreateFile(tx, poc.ID, "point_of_contacts", *fileInfo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		// delete file from s3
		cpoc.AWSProvider.DeleteFile(fileInfo.FileURL)

		return nil, &errors.FailedToCreatePointOfContact
	}

	return poc, nil
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

	return transactions.GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
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

	pointOfContact, err := transactions.GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
	if err != nil {
		return err
	}	

	if err := cpoc.AWSProvider.DeleteFile(pointOfContact.PhotoFile.FileURL); err != nil {
		return err
	}

	tx := cpoc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return &errors.FailedToDeleteClubPointOfContact
	}

	// delete file
	err = transactions.DeleteFile(tx, pointOfContact.PhotoFile.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = transactions.DeleteClubPointOfContact(tx, *clubIdAsUUID, *pocIdAsUUID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return &errors.FailedToDeleteClubPointOfContact
	}

	return transactions.DeleteClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
}