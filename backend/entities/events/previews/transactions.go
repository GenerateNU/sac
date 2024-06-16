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
	for index, event := range events {
		eventIDs[index] = event.ID
	}

	var eventTags []struct {
		EventID uuid.UUID `gorm:"column:event_id"`
		TagID   uuid.UUID `gorm:"column:tag_id"`
		Name    string    `gorm:"column:name"`
	}
	if err := db.Table("tags").Select("tags.id as tag_id, tags.name, event_tags.event_id").
		Joins("JOIN event_tags ON event_tags.tag_id = tags.id").
		Where("event_tags.event_id IN ?", eventIDs).
		Scan(&eventTags).Error; err != nil {
		return nil, err
	}

	eventTagsMap := make(map[uuid.UUID][]models.Tag)
	for _, et := range eventTags {
		eventTagsMap[et.EventID] = append(eventTagsMap[et.EventID],
			models.Tag{
				Model: models.Model{
					ID: et.TagID,
				},
				Name: et.Name,
			})
	}

	var hosts []models.Club
	if err := db.Joins("JOIN events ON events.host = clubs.id").Where("events.id IN ?", eventIDs).Find(&hosts).Error; err != nil {
		return nil, err
	}

	hostsMap := make(map[uuid.UUID]models.Club)
	for _, host := range hosts {
		hostsMap[host.ID] = host
	}

	eventPreviews := make([]EventPreview, len(events))
	for index, event := range events {
		eventPreviews[index] = *event.Into(eventTagsMap[event.ID], hostsMap[event.Host].Name, hostsMap[event.Host].Logo)
	}

	return eventPreviews, nil
}
