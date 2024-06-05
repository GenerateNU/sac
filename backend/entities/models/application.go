package models

import "github.com/google/uuid"

type Application struct {
	Model

	Title string `json:"title"`
	Link  string `json:"link"`

	RecruitmentID *uuid.UUID `json:"-"`
}
