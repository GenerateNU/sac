package series

import (
	"fmt"
	"time"

	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/GenerateNU/sac/backend/protos"
	"github.com/GenerateNU/sac/backend/utilities"
)

type CreateEventSeriesRequestBody struct {
	models.CreateEventRequestBody
	Recurrence  protos.StringRecurrence `json:"recurrence" validate:"required"`
	Termination time.Time               `json:"termination" validate:"required"`
}

func (c *CreateEventSeriesRequestBody) Into() (*ValidatedCreateEventSeriesRequestBody, error) {
	errors := make(map[string]string)
	recurrence, err := protos.RecurrenceFrom(c.Recurrence)
	if err != nil {
		errors["Recurrence"] = fmt.Sprintf("%s failed %s", "Recurrence", "StringRecurrenceProto")
	}

	if len(errors) > 0 {
		return nil, utilities.InvalidRequestData(errors)
	}

	return &ValidatedCreateEventSeriesRequestBody{
		CreateEventRequestBody: c.CreateEventRequestBody,
		Recurrence:             *recurrence,
		Termination:            c.Termination,
	}, nil
}

type ValidatedCreateEventSeriesRequestBody struct {
	models.CreateEventRequestBody
	Recurrence  protos.Recurrence `json:"recurrence"`
	Termination time.Time         `json:"termination"`
}
