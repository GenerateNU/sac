package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]models.User, *errors.Error) {
	var users []models.User

	if err := db.Omit("password_hash").Find(&users).Error; err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetAllUsers}
	}

	return users, nil
}

func GetUser(db *gorm.DB, id uint) (*models.User, *errors.Error) {
	var user models.User
	if err := db.Omit("password_hash").First(&user, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.UserNotFound}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetUser}
		}
	}

	return &user, nil
}

func CreateUser(db *gorm.DB, user *models.User) (*models.User, *errors.Error) {
	if err := db.Create(user).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, &errors.Error{StatusCode: fiber.StatusBadRequest, Message: errors.UserAlreadyExists}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToCreateUser}
		}
	}

	return user, nil
}

func UpdateUser(db *gorm.DB, id uint, user models.User) (*models.User, *errors.Error) {
	var existingUser models.User

	err := db.First(&existingUser, id).Error
	if err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.UserNotFound}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetUser}
		}
	}

	if err := db.Model(&existingUser).Updates(&user).Error; err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToUpdateUser}
	}

	return &existingUser, nil
}

func DeleteUser(db *gorm.DB, id uint) *errors.Error {
	var deletedUser models.User

	result := db.Where("id = ?", id).Delete(&deletedUser)
	if result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.UserNotFound}
		} else {
			return &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToDeleteUser}
		}
	}
	return nil
}
