package oauth

import (
	"errors"
	"time"

	"github.com/GenerateNU/sac/backend/auth"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetRefreshToken(db *gorm.DB, userID *uuid.UUID) (*string, error) {
	var refreshToken string

	if err := db.Model(&models.OAuthToken{}).Where("user_id = ?", userID).Select("refresh_token").Find(&refreshToken).Error; err != nil {
		return nil, err
	}

	if refreshToken == "" {
		return nil, errors.New("refresh token not found")
	}

	return &refreshToken, nil
}

func GetCSRFToken(db *gorm.DB, userID *uuid.UUID) (*string, error) {
	var csrfToken string

	if err := db.Model(&models.OAuthToken{}).Where("user_id = ?", userID).Select("csrf_token").Find(&csrfToken).Error; err != nil {
		return nil, err
	}

	if csrfToken == "" {
		return nil, errors.New("csrf token not found")
	}

	return &csrfToken, nil
}

func SaveRefreshToken(db *gorm.DB, userID *uuid.UUID, refreshToken string) error {
	if refreshToken == "" {
		return errors.New("refresh token is required")
	}

	calendarToken := models.OAuthToken{
		UserID: 	 *userID,
		RefreshToken: refreshToken,
	}

	if err := db.Create(&calendarToken).Error; err != nil {
		return err
	}

	return nil
}

func CreateCalendarToken(db *gorm.DB, jwtClient auth.JWTClientInterface, userID *uuid.UUID) (*string, error) {
	// Generate CSRF token:
	csrfToken, err := jwtClient.GenerateToken(auth.Claims{
		StandardClaims: &jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			Issuer:   userID.String(),
		}},
		"csrf",
	)
	if err != nil {
		return nil, err
	}

	csrfTokenAsString := string(csrfToken)
	
	// Create a calendar token:
	calendarToken := models.OAuthToken{
		UserID: 	 *userID,
		CSRFToken: 	 csrfTokenAsString,
	}

	if err := db.Create(&calendarToken).Error; err != nil {
		return nil, err
	}

	return &csrfTokenAsString, nil
}

func DeleteCalendarToken(db *gorm.DB, userID *uuid.UUID) error {
	if err := db.Where("user_id = ?", userID).Delete(&models.OAuthToken{}).Error; err != nil {
		return err
	}

	return nil
}
