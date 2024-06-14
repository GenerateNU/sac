package rsvps

import (
	"errors"

	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/clubs"
	"github.com/GenerateNU/sac/backend/entities/events"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/entities/users"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateEventRSVP(db *gorm.DB, eventID uuid.UUID, userID uuid.UUID) error {
	db = cache.SetUseCache(db, true)

	event, err := events.GetEvent(db, eventID)
	if err != nil {
		return err
	}

	if event.IsArchived {
		return utilities.ForbiddenReason(errors.New("cannot RSVP to an archived event"))
	}

	if event.IsDraft {
		return utilities.ForbiddenReason(errors.New("cannot RSVP to a draft event"))
	}

	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	if !event.IsPublic {
		isMember, err := clubs.IsMember(db, *event.Host, userID)
		if err != nil {
			return err
		}

		if !isMember {
			return utilities.ForbiddenReason(errors.New("cannot RSVP to a private event when not a member"))
		}
	}

	if err := db.Model(&user).Association("RSVP").Append(event); err != nil {
		return err
	}

	return nil
}

func GetEventRSVPs(db *gorm.DB, eventID uuid.UUID, pageInfo fiberpaginate.PageInfo) ([]models.User, error) {
	db = cache.SetUseCache(db, true)

	event, err := events.GetEvent(db, eventID)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := db.Scopes(utilities.IntoScope(pageInfo, db)).Model(&event).Association("RSVP").Find(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func DeleteEventRSVP(db *gorm.DB, eventID uuid.UUID, userID uuid.UUID) error {
	db = cache.SetUseCache(db, true)

	user, err := users.GetUser(db, userID)
	if err != nil {
		return err
	}

	event, err := events.GetEvent(db, eventID)
	if err != nil {
		return err
	}

	if err := db.Model(&user).Association("RSVP").Delete(event); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utilities.BadRequest(errors.New("user was never RSVPed to this event"))
		}
		return err
	}

	return nil
}
