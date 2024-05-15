package base

import (
	"errors"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/garrettladley/fiberpaginate"

	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) (*models.User, error) {
	if err := db.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, utilities.ErrDuplicate
		}
		return nil, err
	}

	return user, nil
}

func GetUsers(db *gorm.DB, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	var users []models.User
	if err := db.Omit("password_hash").Scopes(utilities.IntoScope(pageInfo, db)).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserPasswordHash(db *gorm.DB, id uuid.UUID) (*string, error) {
	var user models.User
	if err := db.Select("password_hash").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	return &user.PasswordHash, nil
}

func UpdateUser(db *gorm.DB, id uuid.UUID, user models.User) (*models.User, error) {
	var existingUser models.User

	err := db.First(&existingUser, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilities.ErrNotFound
		}
		return nil, err
	}

	if err := db.Model(&existingUser).Updates(&user).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func DeleteUser(db *gorm.DB, id uuid.UUID) error {
	if result := db.Delete(&models.User{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return utilities.ErrNotFound
		}
		return result.Error
	}
	return nil
}
