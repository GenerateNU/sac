package base

import (
	"errors"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateOAuthToken(db *gorm.DB, oauthToken models.UserOAuthTokens) error {
	if err := db.Where("user_id = ? AND resource_type = ?", oauthToken.UserID, oauthToken.ResourceType).Save(&oauthToken).Error; err != nil {
		return err
	}

	return nil
}

func GetOAuthToken(db *gorm.DB, userID uuid.UUID, resourceType models.OAuthResource) (*models.UserOAuthTokens, error) {
	var oauthToken models.UserOAuthTokens

	if err := db.Where("user_id = ? AND resource_type = ?", userID, resourceType).First(&oauthToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &oauthToken, nil
}

func SaveOAuthTokens(db *gorm.DB, userID uuid.UUID, tokens models.OAuthToken, resourceType models.OAuthResource) error {
	oauthToken := models.UserOAuthTokens{
		UserID:       userID,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(tokens.ExpiresIn) * time.Second),
		ResourceType: resourceType,
	}

	if err := db.Where("user_id = ? AND resource_type = ?", userID, resourceType).Save(&oauthToken).Error; err != nil {
		return err
	}

	return nil
}

func DeleteOAuthToken(db *gorm.DB, userID uuid.UUID, resourceType models.OAuthResource) error {
	if err := db.Where("user_id = ? AND resource_type = ?", userID, resourceType).Delete(&models.UserOAuthTokens{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utilities.ErrNotFound
		}
		return err
	}

	return nil
}
