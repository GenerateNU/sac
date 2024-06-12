package tags

import (
	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/events"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/tags"
	"github.com/GenerateNU/sac/backend/transactions"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

func CreateEventTags(db *gorm.DB, id uuid.UUID, tags []models.Tag) ([]models.Tag, error) {
	event, err := events.GetEvent(db, id, transactions.PreloadTag())
	if err != nil {
		return nil, err
	}

	if err := db.Model(&event).Association("Tag").Append(tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func GetEventTags(db *gorm.DB, id uuid.UUID) ([]models.Tag, error) {
	db = cache.SetCache(db, true)

	var tags []models.Tag
	event, err := events.GetEvent(db, id, transactions.PreloadTag())
	if err != nil {
		return nil, err
	}

	if err := db.Model(&event).Association("Tag").Find(&tags); err != nil {
		return nil, err
	}
	return tags, nil
}

func DeleteEventTag(db *gorm.DB, id uuid.UUID, tagID uuid.UUID) error {
	event, err := events.GetEvent(db, id, transactions.PreloadTag())
	if err != nil {
		return err
	}

	tag, err := tags.GetTag(db, tagID)
	if err != nil {
		return err
	}

	if err := db.Model(&event).Association("Tag").Delete(&tag); err != nil {
		return err
	}
	return nil
}
