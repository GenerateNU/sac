package previews

import (
	"time"

	"github.com/garrettladley/fiberpaginate"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/events/tags"
	"github.com/GenerateNU/sac/backend/utilities"

	"gorm.io/gorm"
)

func GetEventPreivews(db *gorm.DB, pageInfo *fiberpaginate.PageInfo) ([]EventPreview, error) {
	db = cache.SetUseCache(db, true)

	var events []EventStub
	if err := db.Scopes(utilities.IntoScope(*pageInfo, db)).Find(&events).Error; err != nil {
		return nil, err
	}

	return processEventStubs(db, events)
}

func GetEventPreviewsByTime(db *gorm.DB, pageInfo *fiberpaginate.PageInfo, startTime *time.Time, endTime *time.Time) ([]EventPreview, error) {
	db = cache.SetUseCache(db, true)

	var events []EventStub
	query := db.Model(&events)

	if startTime != nil && endTime != nil {
		query = query.Where("start_time >= ? AND end_time <= ?", *startTime, *endTime)
	}

	if pageInfo != nil {
		query = query.Scopes(utilities.IntoScope(*pageInfo, db))
	}

	if err := query.Find(&events).Error; err != nil {
		return nil, err
	}

	return processEventStubs(db, events)
}

func processEventStubs(db *gorm.DB, events []EventStub) ([]EventPreview, error) {
	eventPreviews := make([]EventPreview, len(events))
	for index, event := range events {
		eventTags, err := tags.GetEventTags(db, event.ID)
		if err != nil {
			return nil, err
		}

		host, err := clubs.GetClub(db, event.Host)
		if err != nil {
			return nil, err
		}

		eventPreviews[index] = *event.Into(eventTags, host.Name, host.Logo)
	}
	return eventPreviews, nil
}
