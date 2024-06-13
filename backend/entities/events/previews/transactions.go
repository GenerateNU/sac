package previews

import (
	"time"

	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/utilities"

	"gorm.io/gorm"
)

func GetEventPreviews(db *gorm.DB, pageInfo *fiberpaginate.PageInfo) ([]EventPreview, error) {
	db = cache.SetUseCache(db, true)

	var events []EventStub
	if err := db.Scopes(utilities.IntoScope(*pageInfo, db)).Find(&events).Error; err != nil {
		return nil, err
	}

	return processEventStubs(db, events)
}

func GetEventPreviewsByTime(db *gorm.DB, pageInfo *fiberpaginate.PageInfo, startTime, endTime *time.Time) ([]EventPreview, error) {
	db = cache.SetUseCache(db, true)

	var events []EventStub
	query := db.Model(&events).Scopes(utilities.IntoScope(*pageInfo, db))

	if startTime != nil && endTime != nil {
		query = query.Where("start_time >= ? AND end_time <= ?", *startTime, *endTime)
	}

	if err := query.Find(&events).Error; err != nil {
		return nil, err
	}

	return processEventStubs(db, events)
}

func processEventStubs(db *gorm.DB, events []EventStub) ([]EventPreview, error) {
	eventIDs := make([]uuid.UUID, len(events))
	for i, event := range events {
		eventIDs[i] = event.ID
	}

	var eventTagsMap map[uuid.UUID][]models.Tag
	if err := db.Model(&models.Tag{}).Joins("JOIN event_tags ON event_tags.tag_id = tags.id").
		Where("event_tags.event_id IN ?", eventIDs).
		Find(&eventTagsMap).Error; err != nil {
		return nil, err
	}

	var hostsMap map[uuid.UUID]models.Club
	if err := db.Model(&models.Club{}).Where("id IN ?", eventIDs).Find(&hostsMap).Error; err != nil {
		return nil, err
	}

	eventPreviews := make([]EventPreview, len(events))
	for i, event := range events {
		eventPreviews[i] = *event.Into(eventTagsMap[event.ID], hostsMap[event.Host].Name, hostsMap[event.Host].Logo)
	}

	return eventPreviews, nil
}
