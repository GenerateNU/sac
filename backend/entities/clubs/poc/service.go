package poc

import (
	"mime/multipart"

	"github.com/GenerateNU/sac/backend/entities/files/base"
	"github.com/GenerateNU/sac/backend/entities/models"

	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubPointOfContactServiceInterface interface {
	GetClubPointOfContacts(clubID string) ([]models.PointOfContact, error)
	GetClubPointOfContact(clubID, pocID string) (*models.PointOfContact, error)
	CreateClubPointOfContact(clubID string, pointOfContactBody models.CreatePointOfContactBody, fileHeader *multipart.FileHeader) (*models.PointOfContact, error)
	UpdateClubPointOfContactPhoto(clubID, pocID string, fileHeader *multipart.FileHeader) (*models.PointOfContact, error)
	UpdateClubPointOfContact(clubID, pocID string, pointOfContactBody models.UpdatePointOfContactBody) (*models.PointOfContact, error)
	DeleteClubPointOfContact(clubID, pocID string) error
}

type ClubPointOfContactService struct {
	types.ServiceParams
}

func NewClubPointOfContactService(serviceParams types.ServiceParams) ClubPointOfContactServiceInterface {
	return &ClubPointOfContactService{serviceParams}
}

func (cpoc *ClubPointOfContactService) GetClubPointOfContacts(clubID string) ([]models.PointOfContact, error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubPointOfContacts(cpoc.DB, *clubIdAsUUID)
}

func (cpoc *ClubPointOfContactService) GetClubPointOfContact(clubID, pocID string) (*models.PointOfContact, error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, err
	}

	return GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
}

func (cpoc *ClubPointOfContactService) CreateClubPointOfContact(clubID string, pointOfContactBody models.CreatePointOfContactBody, fileHeader *multipart.FileHeader) (*models.PointOfContact, error) {
	if err := utilities.Validate(cpoc.Validate, pointOfContactBody); err != nil {
		return nil, err
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	_, err = GetClubPointOfContactByClubIDAndEmail(cpoc.DB, *clubIdAsUUID, pointOfContactBody.Email)
	if err == nil {
		return nil, err
	}

	fileInfo, err := cpoc.Integrations.File.UploadFile("point_of_contacts", fileHeader, []file.FileType{file.IMAGE})
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
		if err := cpoc.Integrations.File.DeleteFile(fileInfo.FileURL); err != nil {
			return nil, err
		}

		return nil, err
	}

	poc.PhotoFile = *file

	return poc, nil
}

func (cpoc *ClubPointOfContactService) UpdateClubPointOfContactPhoto(clubID, pocID string, fileHeader *multipart.FileHeader) (*models.PointOfContact, error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, err
	}

	pointOfContact, err := GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
	if err != nil {
		return nil, err
	}

	if err := cpoc.Integrations.File.DeleteFile(pointOfContact.PhotoFile.FileURL); err != nil {
		return nil, err
	}

	fileInfo, err := cpoc.Integrations.File.UploadFile("point_of_contacts", fileHeader, []file.FileType{file.IMAGE})
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

func (cpoc *ClubPointOfContactService) UpdateClubPointOfContact(clubID, pocID string, pointOfContactBody models.UpdatePointOfContactBody) (*models.PointOfContact, error) {
	if err := utilities.Validate(cpoc.Validate, pointOfContactBody); err != nil {
		return nil, err
	}

	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return nil, err
	}

	return UpdateClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID, pointOfContactBody)
}

func (cpoc *ClubPointOfContactService) DeleteClubPointOfContact(clubID, pocID string) error {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return err
	}

	pocIdAsUUID, err := utilities.ValidateID(pocID)
	if err != nil {
		return err
	}

	pointOfContact, err := GetClubPointOfContact(cpoc.DB, *clubIdAsUUID, *pocIdAsUUID)
	if err != nil {
		return err
	}

	if err := cpoc.Integrations.File.DeleteFile(pointOfContact.PhotoFile.FileURL); err != nil {
		return err
	}

	tx := cpoc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
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
		return err
	}

	return nil
}
