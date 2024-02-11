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

func GetUsers(db *gorm.DB, limit int, offset int) ([]models.User, *errors.Error) {
	var users []models.User

	if err := db.Omit("password_hash").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, &errors.FailedToGetUsers
	}

	return users, nil
}

func GetUser(db *gorm.DB, id uuid.UUID) (*models.User, *errors.Error) {
	var user models.User
	if err := db.Omit("password_hash").First(&user, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToGetUser
		}
	}

	return &user, nil
}

func GetUserWithFollowers(db *gorm.DB, id uuid.UUID) (*models.User, *errors.Error) {
	var user models.User
	if err := db.Preload("Follower").Omit("password_hash").First(&user, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.UserNotFound
		} else {
			return nil, &errors.FailedToGetUser
		}
	}

	return &user, nil
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

func DeleteUser(db *gorm.DB, id uuid.UUID) *errors.Error {
	result := db.Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.UserNotFound
		} else {
			return &errors.FailedToDeleteUser
		}
	}
	return nil
}
