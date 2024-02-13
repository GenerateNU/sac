package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserFollowerServiceInterface interface {
	CreateFollowing(userId string, clubId string) *errors.Error
	DeleteFollowing(userId string, clubId string) *errors.Error
	GetFollowing(userId string) ([]models.Club, *errors.Error)
}

type UserFollowerService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewUserFollowerService(db *gorm.DB, validate *validator.Validate) *UserFollowerService {
	return &UserFollowerService{DB: db, Validate: validate}
}

func (u *UserFollowerService) CreateFollowing(userId string, clubId string) *errors.Error {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return err
	}
	clubIdAsUUID, err := utilities.ValidateID(clubId)
	if err != nil {
		return err
	}
	return transactions.CreateFollowing(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserFollowerService) DeleteFollowing(userId string, clubId string) *errors.Error {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return err
	}
	clubIdAsUUID, err := utilities.ValidateID(clubId)
	if err != nil {
		return err
	}
	return transactions.DeleteFollowing(u.DB, *userIdAsUUID, *clubIdAsUUID)
}

func (u *UserFollowerService) GetFollowing(userId string) ([]models.Club, *errors.Error) {
	userIdAsUUID, err := utilities.ValidateID(userId)
	if err != nil {
		return nil, err
	}

	return transactions.GetClubFollowing(u.DB, *userIdAsUUID)
}
