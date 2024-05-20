package users

import (
	"context"
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id uuid.UUID, preloads ...transactions.OptionalQuery) (*models.User, error) {
	var user models.User

	query := db

	for _, preload := range preloads {
		query = preload(query)
	}

	if err := query.Omit("password_hash").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func UpdateEmailVerification(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	result := db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Update("is_verified", true)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}
	return nil
}

func UpdatePassword(db *gorm.DB, id uuid.UUID, passwordHash string) error {
	result := db.Model(&models.User{}).Where("id = ?", id).Update("password_hash", passwordHash)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}
	return nil
}
