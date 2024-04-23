package poc

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/files/base"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/file"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubPointOfContactServiceInterface interface {
	GetClubPointOfContacts(clubID string) ([]models.PointOfContact, *errors.Error)
	GetClubPointOfContact(clubID, pocID string) (*models.PointOfContact, *errors.Error)
	CreateClubPointOfContact(clubID string, pointOfContactBody models.CreatePointOfContactBody, fileHeader *multipart.FileHeader) (*models.PointOfContact, *errors.Error)
	UpdateClubPointOfContactPhoto(clubID, pocID string, fileHeader *multipart.FileHeader) (*models.PointOfContact, *errors.Error)
	UpdateClubPointOfContact(clubID, pocID string, pointOfContactBody models.UpdatePointOfContactBody) (*models.PointOfContact, *errors.Error)
	DeleteClubPointOfContact(clubID, pocID string) *errors.Error
}

type ClubPointOfContactService struct {
	types.ServiceParams
}

func NewClubPointOfContactService(serviceParams types.ServiceParams) ClubPointOfContactServiceInterface {
	return &ClubPointOfContactService{serviceParams}
}

func (cpoc *ClubPointOfContactService) GetClubPointOfContacts(clubID string) ([]models.PointOfContact, *errors.Error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	return GetClubPointOfContacts(cpoc.DB, *clubIdAsUUID)
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

	return GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
}

func (cpoc *ClubPointOfContactService) CreateClubPointOfContact(clubID string, pointOfContactBody models.CreatePointOfContactBody, fileHeader *multipart.FileHeader) (*models.PointOfContact, *errors.Error) {
	if err := cpoc.Validate.Struct(pointOfContactBody); err != nil {
		return nil, &errors.FailedToValidatePointOfContact
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	_, err = GetClubPointOfContactByClubIDAndEmail(cpoc.DB, *clubIdAsUUID, pointOfContactBody.Email)
	if err == nil {
		return nil, &errors.PointOfContactAlreadyExists
	}

	fileInfo, err := cpoc.Deps.File.UploadFile("point_of_contacts", fileHeader, []file.FileType{file.IMAGE})
	if err != nil {
		return nil, err
	}

	tx := cpoc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	poc, err := CreateClubPointOfContact(tx, *clubIdAsUUID, pointOfContactBody)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	file, err := base.CreateFile(tx, poc.ID, "point_of_contacts", *fileInfo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		if err := cpoc.Deps.File.DeleteFile(fileInfo.FileURL); err != nil {
			return nil, err
		}

		return nil, &errors.FailedToCreatePointOfContact
	}

	poc.PhotoFile = *file

	return poc, nil
}

func (cpoc *ClubPointOfContactService) UpdateClubPointOfContactPhoto(clubID, pocID string, fileHeader *multipart.FileHeader) (*models.PointOfContact, *errors.Error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, &errors.FailedToValidatePointOfContactId
	}

	pointOfContact, err := GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
	if err != nil {
		return nil, err
	}

	if err := cpoc.Deps.File.DeleteFile(pointOfContact.PhotoFile.FileURL); err != nil {
		return nil, err
	}

	fileInfo, err := cpoc.Deps.File.UploadFile("point_of_contacts", fileHeader, []file.FileType{file.IMAGE})
	if err != nil {
		return nil, err
	}

	file, err := base.UpdateFile(cpoc.DB, pointOfContact.PhotoFile.ID, *fileInfo)
	if err != nil {
		return nil, err
	}

	pointOfContact.PhotoFile = *file

	return pointOfContact, nil
}

func (cpoc *ClubPointOfContactService) UpdateClubPointOfContact(clubID, pocID string, pointOfContactBody models.UpdatePointOfContactBody) (*models.PointOfContact, *errors.Error) {
	if err := cpoc.Validate.Struct(pointOfContactBody); err != nil {
		return nil, &errors.FailedToValidatePointOfContact
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, &errors.FailedToValidatePointOfContactId
	}

	return UpdateClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID, pointOfContactBody)
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

	pointOfContact, err := GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
	if err != nil {
		return err
	}

	if err := cpoc.Deps.File.DeleteFile(pointOfContact.PhotoFile.FileURL); err != nil {
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

	err = base.DeleteFile(tx, pointOfContact.PhotoFile.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = DeleteClubPointOfContact(tx, *clubIdAsUUID, *pocIdAsUUID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return &errors.FailedToDeleteClubPointOfContact
	}

	return nil
}
