package categories

type CategoryRequestBody struct {
	Name string `json:"name" validate:"required,max=255"`
}
