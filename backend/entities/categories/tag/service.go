package tag

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/middleware"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type CategoryTagServiceInterface interface {
	GetTagsByCategory(categoryID string, pagination middleware.Pagination) ([]models.Tag, error)
	GetTagByCategory(categoryID string, tagID string) (*models.Tag, error)
}

type CategoryTagService struct {
	types.ServiceParams
}

func NewCategoryTagService(params types.ServiceParams) CategoryTagServiceInterface {
	return &CategoryTagService{params}
}

func (t *CategoryTagService) GetTagsByCategory(categoryID string, pagination middleware.Pagination) ([]models.Tag, error) {
	categoryIDAsUUID, err := utilities.ValidateID(categoryID)
	if err != nil {
		return nil, err
	}

	return GetTagsByCategory(t.DB, *categoryIDAsUUID, pagination)
}

func (t *CategoryTagService) GetTagByCategory(categoryID string, tagID string) (*models.Tag, error) {
	categoryIDAsUUID, idErr := utilities.ValidateID(categoryID)

	if idErr != nil {
		return nil, idErr
	}

	tagIDAsUUID, idErr := utilities.ValidateID(tagID)

	if idErr != nil {
		return nil, idErr
	}

	return GetTagByCategory(t.DB, *categoryIDAsUUID, *tagIDAsUUID)
}
