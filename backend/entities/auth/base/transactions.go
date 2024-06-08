package base

import (
	"context"
	"errors"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindOrCreateUser(ctx context.Context, db *gorm.DB, user oauth.User) (*models.User, error) {
	var sacUser models.User
	if err := db.WithContext(ctx).Where("email = ?", user.Email).First(&sacUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newUser := &models.User{Email: user.Email, Name: user.Name, Role: models.Student}
			if err := db.WithContext(ctx).Create(newUser).Error; err != nil {
				return newUser, err
			}
		} else {
			return nil, err
		}
	}

	return &sacUser, nil
}

func GetToken(db *gorm.DB, token string) (*models.Verification, error) {
	tokenModel := models.Verification{}
	if err := db.Where("token = ?", token).First(&tokenModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &tokenModel, nil
}

func GetActiveTokenByUserID(db *gorm.DB, userID uuid.UUID) (*models.Verification, error) {
	token := models.Verification{}
	if err := db.Where("user_id = ? AND expires_at > ?", userID, time.Now().UTC()).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &token, nil
}

func DeleteToken(db *gorm.DB, token string) error {
	if err := db.Where("token = ?", token).Delete(&models.Verification{}).Error; err != nil {
		return err
	}

	return nil
}

func SaveToken(db *gorm.DB, userID uuid.UUID, token string, expiry time.Time) error {
	tokenModel := models.Verification{
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiry,
	}

	if err := db.Create(&tokenModel).Error; err != nil {
		return err
	}

	return nil
}
