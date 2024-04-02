package services

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/transactions"
	"github.com/GenerateNU/sac/backend/src/types"
	"github.com/GenerateNU/sac/backend/src/utilities"
)

type CategoryTagServiceInterface interface {
	GetTagsByCategory(categoryID string, limit string, page string) ([]models.Tag, *errors.Error)
	GetTagByCategory(categoryID string, tagID string) (*models.Tag, *errors.Error)
}

type CategoryTagService struct {
	types.ServiceParams
}

func NewCategoryTagService(params types.ServiceParams) *CategoryTagService {
	return &CategoryTagService{params}
}

func (t *CategoryTagService) GetTagsByCategory(categoryID string, limit string, page string) ([]models.Tag, *errors.Error) {
	categoryIDAsUUID, err := utilities.ValidateID(categoryID)
	if err != nil {
		return nil, err
	}

	limitAsInt, err := utilities.ValidateNonNegative(limit)
	if err != nil {
		return nil, &errors.FailedToValidateLimit
	}

	pageAsInt, err := utilities.ValidateNonNegative(page)
	if err != nil {
		return nil, &errors.FailedToValidatePage
	}

	return transactions.GetTagsByCategory(t.DB, *categoryIDAsUUID, *limitAsInt, *pageAsInt)
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
