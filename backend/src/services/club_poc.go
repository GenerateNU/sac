package services

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/file"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

	// pocExists, err := transactions.ClubPointOfContactExists(cpoc.DB, *clubIdAsUUID, pointOfContactBody.Email)
	// if err != nil {
	// 	return nil, err
	// }

	fileInfo, err := cpoc.AWSProvider.UploadFile("club_point_of_contact", fileHeader)
	if err != nil {
		return nil, err
	}

	tx := cpoc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. once upload is successful, create file in db
	file, err := transactions.CreateFile(tx, uuid.Nil, "club_point_of_contact", *fileInfo)
	if err != nil {
		return nil, err
	}

	// 2. create point of contact
	poc, err := transactions.CreateClubPointOfContact(tx, *clubIdAsUUID, pointOfContactBody, file.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 3. update file with point of contact id
	if err := tx.Model(&file).Update("owner_id", poc.ID).Error; err != nil {
		tx.Rollback()
		return nil, &errors.FailedToUpdateFile
	}

	tx.Commit()
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

	if pointOfContact.PhotoFileID != uuid.Nil {
		if err := cpoc.AWSProvider.DeleteFile(pointOfContact.PhotoFileID.String()); err != nil {
			return err
		}
	} else {
		return &errors.FailedToGetFile
	}

	return transactions.DeleteClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
}