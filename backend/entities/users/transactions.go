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

	if err := query.First(&user, id).Error; err != nil {
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

func VerifyEmail(ctx context.Context, db *gorm.DB, token string, userID uuid.UUID) error {
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Model(&models.User{}).Where("id = ?", userID).Update("is_verified", "verified")
	if result.RowsAffected == 0 {
		if result.Error == nil {
			tx.Rollback()
			return utilities.ErrNotFound
		}
		tx.Rollback()
		return result.Error
	}

	if err := tx.Where("token = ? AND user_id = ?", token, userID).Delete(&models.Verification{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
