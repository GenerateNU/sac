package tags

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type CategoryTagServiceInterface interface {
	GetTagsByCategory(categoryID string, pageInfo fiberpaginate.PageInfo) ([]models.Tag, error)
	GetTagByCategory(categoryID string, tagID string) (*models.Tag, error)
}

type CategoryTagService struct {
	types.ServiceParams
}

func NewCategoryTagService(params types.ServiceParams) CategoryTagServiceInterface {
	return &CategoryTagService{params}
}

func (t *CategoryTagService) GetTagsByCategory(categoryID string, pageInfo fiberpaginate.PageInfo) ([]models.Tag, error) {
	categoryIDAsUUID, err := utilities.ValidateID(categoryID)
	if err != nil {
		return nil, err
	}

	return GetTagsByCategory(t.DB, *categoryIDAsUUID, pageInfo)
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
