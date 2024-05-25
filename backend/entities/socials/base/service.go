package base

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type SocialServiceInterface interface {
	GetSocials(pageInfo fiberpaginate.PageInfo) ([]models.Social, error)
	GetSocial(socialID string) (*models.Social, error)
	DeleteSocial(socialID string) error
}

type SocialService struct {
	types.ServiceParams
}

func NewSocialService(serviceParams types.ServiceParams) SocialServiceInterface {
	return &SocialService{serviceParams}
}

func (c *SocialService) GetSocials(pageInfo fiberpaginate.PageInfo) ([]models.Social, error) {
	return GetSocials(c.DB, pageInfo)
}

func (c *SocialService) GetSocial(socialID string) (*models.Social, error) {
	idAsUUID, err := utilities.ValidateID(socialID)
	if err != nil {
		return nil, err
	}

	return GetSocial(c.DB, *idAsUUID)
}

func (c *SocialService) DeleteSocial(socialID string) error {
	idAsUUID, err := utilities.ValidateID(socialID)
	if err != nil {
		return err
	}

	return DeleteSocial(c.DB, *idAsUUID)
}
