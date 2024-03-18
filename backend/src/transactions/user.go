package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) (*models.User, *errors.Error) {
	if err := db.Create(user).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &errors.UserAlreadyExists
		} else {
			return nil, &errors.FailedToCreateUser
		}
	}

	return user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*models.User, *errors.Error) {
	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, &errors.UserNotFound
	}

	return &user, nil
}

func GetUsers(db *gorm.DB, limit int, page int) ([]models.User, *errors.Error) {
	var users []models.User

	offset := (page - 1) * limit

	if err := db.Omit("password_hash").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, &errors.FailedToGetUsers
	}

	return users, nil
}

func GetUser(db *gorm.DB, id uuid.UUID, preloads ...OptionalQuery) (*models.User, *errors.Error) {
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

func GetUserPasswordHash(db *gorm.DB, id uuid.UUID) (*string, *errors.Error) {
	var user models.User
	if err := db.Select("password_hash").First(&user, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToGetUser
		}
	}

	return &user.PasswordHash, nil
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

func UpdateUser(db *gorm.DB, id uuid.UUID, user models.User) (*models.User, *errors.Error) {
	var existingUser models.User

	err := db.First(&existingUser, id).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToUpdateTag
		}
	}

	if err := db.Model(&existingUser).Updates(&user).Error; err != nil {
		return nil, &errors.FailedToUpdateUser
	}

	return &existingUser, nil
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

func DeleteUser(db *gorm.DB, id uuid.UUID) *errors.Error {
	if result := db.Delete(&models.User{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.UserNotFound
		} else {
			return &errors.FailedToDeleteUser
		}
	}
	return nil
}
