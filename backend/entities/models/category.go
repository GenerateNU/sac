package models

type Category struct {
	Model

	Name string `json:"name"`
	Tag  []Tag  `json:"-"`
}

type CategoryRequestBody struct {
	Name string `json:"name" validate:"required,max=255"`
}
