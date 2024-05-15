package models

type CalendarTokenRequest struct {
	Code *string `json:"code" validate:"required"`
}