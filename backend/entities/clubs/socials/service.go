package socials

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
)

type ClubSocialServiceInterface interface {
	GetClubSocials(clubID string) ([]models.Social, error)
	PutClubSocial(clubID string, contactBody PutSocialRequestBody) (*models.Social, error)
}

type ClubSocialService struct {
	types.ServiceParams
}

func NewClubSocialService(params types.ServiceParams) ClubSocialServiceInterface {
	return &ClubSocialService{params}
}

func (c *ClubSocialService) GetClubSocials(clubID string) ([]models.Social, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	return GetClubSocials(c.DB, *idAsUUID)
}

func (c *ClubSocialService) PutClubSocial(clubID string, contactBody PutSocialRequestBody) (*models.Social, error) {
	idAsUUID, err := utilities.ValidateID(clubID)
	if err != nil {
		return nil, err
	}

	if err := utilities.Validate(c.Validate, contactBody); err != nil {
		return nil, err
	}

	contact, err := utilities.MapJsonTags(contactBody, &models.Social{})
	if err != nil {
		return nil, err
	}

	contact.ClubID = *idAsUUID

	return PutClubSocial(c.DB, *contact)
}
