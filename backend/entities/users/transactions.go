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

func GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}
