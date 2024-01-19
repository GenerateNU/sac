package types

type UserParams struct {
	NUID      string `json:"nuid" validate:"omitempty,number,len=9"`
	FirstName string `json:"first_name" validate:"omitempty,max=255"`
	LastName  string `json:"last_name" validate:"omitempty,max=255"`
	Email     string `json:"email" validate:"omitempty,email,neu_email"`
	Password  string `json:"password" validate:"omitempty,password"`
	College   string `json:"college" validate:"omitempty,oneof=CAMD DMSB KCCS CE BCHS SL CPS CS CSSH"`
	Year      uint   `json:"year" validate:"omitempty,min=1,max=6"`
}
