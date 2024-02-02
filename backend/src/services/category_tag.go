package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/utilities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CategoryTagServiceInterface interface {
	GetTagByCategory(categoryID string, tagID string) (*models.Tag, *errors.Error)
}

type CategoryTagService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewCategoryTagService(db *gorm.DB, validate *validator.Validate) *CategoryTagService {
	return &CategoryTagService{DB: db, Validate: validate}
}

func (t *CategoryTagService) GetTagByCategory(categoryID string, tagID string) (*models.Tag, *errors.Error) {
	categoryIDAsUUID, idErr := utilities.ValidateID(categoryID)

	if idErr != nil {
		return nil, idErr
	}

	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return nil, idErr
	}

	return transactions.GetTagByCategory(t.DB, *categoryIDAsUUID, *tagIDAsUUID)
}
