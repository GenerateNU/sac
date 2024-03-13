package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ClubPointOfContactServiceInterface interface {
	GetClubPointOfContacts(clubId string) ([]models.PointOfContact, *errors.Error)
	GetClubPointOfContact(clubId, pocID string) (*models.PointOfContact, *errors.Error)
	// UpsertClubPointOfContact(clubId string, pointOfContactBody models.CreatePointOfContactBody) (*models.PointOfContact, *errors.Error)
	DeleteClubPointOfContact(clubId, pocID string) *errors.Error
}

type ClubPointOfContactService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewClubPointOfContactService(db *gorm.DB, validate *validator.Validate) *ClubPointOfContactService {
	return &ClubPointOfContactService{DB: db, Validate: validate}
}

func (cpoc *ClubPointOfContactService) GetClubPointOfContacts(clubID string) ([]models.PointOfContact, *errors.Error) {
	clubIdAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateClub
	}

	return transactions.GetClubPointOfContacts(cpoc.DB, *clubIdAsUUID)
}

// func (cpoc *ClubPointOfContactService) UpsertPointOfContact(clubId string, pointOfContactBody models.CreatePointOfContactBody) (*models.PointOfContact, *errors.Error) {
// 	if err := cpoc.Validate.Struct(pointOfContactBody); err != nil {
// 		return nil, &errors.FailedToValidatePointOfContact
// 	}

// 	pointOfContact, err := utilities.MapRequestToModel(pointOfContactBody, &models.PointOfContact{})
// 	if err != nil {
// 		print(err.Error())
// 		return nil, &errors.FailedToMapRequestToModel
// 	}

// 	clubIdAsUUID, idErr := utilities.ValidateID(clubId)
// 	var file models.File
// 	if pointOfContactBody.PhotoFileID != uuid.Nil {
// 		if err := cpoc.DB.First(&file, "id = ?", pointOfContactBody.PhotoFileID).Error; err != nil {
// 			return nil, &errors.CannotFindFile
// 		}
// 	}
// 	pointOfContact.ClubID = *clubIdAsUUID
// 	if idErr != nil {
// 		return nil, &errors.FailedToValidateClub
// 	}
// 	poc, upsertErr := transactions.UpsertPointOfContact(cpoc.DB, pointOfContact)

// 	if upsertErr != nil {
// 		return poc, upsertErr
// 	}

// 	if pointOfContactBody.PhotoFileID != uuid.Nil {
// 		file.OwnerType = "point_of_contact"
// 		file.OwnerID = poc.ID
// 	}
// 	cpoc.DB.Save(&file)
// 	pointOfContact.PhotoFile = &file
// 	return poc, nil
// }

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

// Delete A Point of Contact
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