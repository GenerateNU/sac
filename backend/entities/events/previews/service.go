package previews

import (
	"errors"

	"github.com/GenerateNU/sac/backend/types"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/garrettladley/fiberpaginate"
)

type Service interface {
	GetEventPreviews(pageInfo *fiberpaginate.PageInfo, start string, end string) ([]EventPreview, error)
}

type Handler struct {
	types.ServiceParams
}

func NewService(serviceParams types.ServiceParams) Service {
	return &Handler{serviceParams}
}

func (h *Handler) GetEventPreviews(pageInfo *fiberpaginate.PageInfo, start string, end string) ([]EventPreview, error) {
	if start == "" || end == "" {
		return GetEventPreviews(h.DB, pageInfo)
	}

	startTime, err := utilities.ParseTime(start, utilities.YYYY_dash_MM_dash_DD)
	if err != nil {
		return nil, utilities.BadRequest(err)
	}

	endTime, err := utilities.ParseTime(end, utilities.YYYY_dash_MM_dash_DD)
	if err != nil {
		return nil, utilities.BadRequest(err)
	}

	if startTime.After(*endTime) {
		return nil, utilities.BadRequest(errors.New("start time must be before end time"))
	}

	return GetEventPreviewsByTime(h.DB, pageInfo, startTime, endTime)
}
