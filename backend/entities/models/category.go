package models

type Category struct {
	Model

	Name string `json:"name"`
	Tag  []Tag  `json:"-"`
}
