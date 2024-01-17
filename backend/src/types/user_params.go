package types

type UserParams struct {
	NUID      string `json:"nuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	College   string `json:"college"`
	Year      uint   `json:"year"`
}
