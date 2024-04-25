package models

type Category struct {
	Model

	Name string `json:"name" validate:"required,max=255"`
	Tag  []Tag  `json:"-" validate:"-"`
}

type CategoryRequestBody struct {
	Name string `json:"name" validate:"required,max=255"`
}
