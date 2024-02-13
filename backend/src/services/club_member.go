package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ClubMemberServiceInterface interface {
	GetClubMembers(clubID string, limit string, page string) ([]models.User, *errors.Error)
}

type ClubMemberService struct {
	DB *gorm.DB
}

func NewClubMemberService(db *gorm.DB, validate *validator.Validate) *ClubMemberService {
	return &ClubMemberService{DB: db}
}

func (cms *ClubMemberService) GetClubMembers(clubID string, limit string, page string) ([]models.User, *errors.Error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, &errors.FailedToValidateID
	}

	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetClubMembers(cms.DB, *idAsUUID, *limitAsInt, *pageAsInt)
}
