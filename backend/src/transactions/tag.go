package transactions

import (
	stdliberrors "errors"

	"github.com/GenerateNU/sac/backend/src/errors"

	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTag(db *gorm.DB, tag models.Tag) (*models.Tag, *errors.Error) {
	if err := db.Create(&tag).Error; err != nil {
		return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToCreateTag}
	}

	return &tag, nil
}

func GetTag(db *gorm.DB, id uint) (*models.Tag, *errors.Error) {
	var tag models.Tag

	if err := db.First(&tag, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.TagNotFound}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToGetTag}
		}
	}

	return &tag, nil
}

func UpdateTag(db *gorm.DB, id uint, tag models.Tag) (*models.Tag, *errors.Error) {
	if err := db.Model(&models.Tag{}).Where("id = ?", id).Updates(tag).First(&tag, id).Error; err != nil {
		if stdliberrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.TagNotFound}
		} else {
			return nil, &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToUpdateTag}
		}
	}

	return &tag, nil

}

func DeleteTag(db *gorm.DB, id uint) *errors.Error {
	if result := db.Delete(&models.Tag{}, id); result.RowsAffected == 0 {
		if result.Error == nil {
			return &errors.Error{StatusCode: fiber.StatusNotFound, Message: errors.TagNotFound}
		} else {
			return &errors.Error{StatusCode: fiber.StatusInternalServerError, Message: errors.FailedToDeleteTag}
		}
	}

	return nil
}
