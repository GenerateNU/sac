package users

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/errors"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id uuid.UUID, preloads ...transactions.OptionalQuery) (*models.User, *errors.Error) {
	var user models.User

	query := db

	for _, preload := range preloads {
		query = preload(query)
	}

	if err := query.Omit("password_hash").First(&user, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToGetUser
		}
	}

	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, *errors.Error) {
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, &errors.UserNotFound
	}

	return &user, nil
}

func UpdateEmailVerification(db *gorm.DB, id uuid.UUID) *errors.Error {
	result := db.Model(&models.User{}).Where("id = ?", id).Update("is_verified", true)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.UserNotFound
		} else {
			return &errors.FailedToUpdateEmailVerification
		}
	}
	return nil
}

func UpdatePassword(db *gorm.DB, id uuid.UUID, passwordHash string) *errors.Error {
	result := db.Model(&models.User{}).Where("id = ?", id).Update("password_hash", passwordHash)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.UserNotFound
		} else {
			return &errors.FailedToUpdateUser
		}
	}
	return nil
}
