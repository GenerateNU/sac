package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ClubServiceInterface interface {
	GetClubs(limit string, page string) ([]models.Club, *errors.Error)
	GetClub(id string) (*models.Club, *errors.Error)
	CreateClub(clubBody models.CreateClubRequestBody) (*models.Club, *errors.Error)
	UpdateClub(id string, clubBody models.UpdateClubRequestBody) (*models.Club, *errors.Error)
	DeleteClub(id string) *errors.Error
	GetClubMembers(clubID string) ([]models.User, *errors.Error)
	CreateMembership(clubID string, userID string) *errors.Error
	CreateMembershipsByEmail(clubID string, emails []string) *errors.Error
	DeleteMembership(clubID string, userID string) *errors.Error
	DeleteMemberships(clubID string, userIDs []string) *errors.Error
	GetUserFollowersForClub(id string) ([]models.User, *errors.Error)
}

type ClubService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewClubService(db *gorm.DB, validate *validator.Validate) *ClubService {
	return &ClubService{DB: db, Validate: validate}
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

func (c *ClubService) GetClubMembers(clubID string) ([]models.User, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	return transactions.GetClubMembers(c.DB, *idAsUUID)
}

func (c *ClubService) CreateMembership(clubID string, userID string) *errors.Error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.CreateMembership(c.DB, *clubIDAsUUID, *userIDAsUUID)
}

func (c *ClubService) CreateMembershipsByEmail(clubID string, emails []string) *errors.Error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.CreateMembershipsByEmail(c.DB, *clubIDAsUUID, emails)
}

func (c *ClubService) DeleteMembership(clubID string, userID string) *errors.Error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	userIDAsUUID, err := utilities.ValidateID(userID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	return transactions.DeleteMembership(c.DB, *clubIDAsUUID, *userIDAsUUID)
}

func (c *ClubService) DeleteMemberships(clubID string, userIDs []string) *errors.Error {
	clubIDAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return &errors.FailedToValidateID
	}

	var userIDsAsUUIDs []uuid.UUID
	for _, id := range userIDs {
		userIDAsUUID, err := utilities.ValidateID(id)
		if err != nil {
			return &errors.FailedToValidateID
		}
		userIDsAsUUIDs = append(userIDsAsUUIDs, *userIDAsUUID)
	}

	return transactions.DeleteMemberships(c.DB, *clubIDAsUUID, userIDsAsUUIDs)
}

func (c *ClubService) GetUserFollowersForClub(id string) ([]models.User, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(id)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}
	return transactions.GetUserFollowersForClub(c.DB, *idAsUUID)
}
