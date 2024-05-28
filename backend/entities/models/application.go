package models

import "github.com/google/uuid"

type Application struct {
	Model

	Title string `json:"title"`
	Link  string `json:"description"`

	RecruitmentID *uuid.UUID `json:"-"`
}
