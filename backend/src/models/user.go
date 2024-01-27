package models

import "github.com/google/uuid"

type UserRole string

const (
	Super     UserRole = "super"
	ClubAdmin UserRole = "clubAdmin"
	Student   UserRole = "student"
)

type College string

const (
	CAMD College = "CAMD" // College of Arts, Media and Design
	DMSB College = "DMSB" // D'Amore-McKim School of Business
	KCCS College = "KCCS" // Khoury College of Computer Sciences
	CE   College = "CE"   // College of Engineering
	BCHS College = "BCHS" // Bouvé College of Health Sciences
	SL   College = "SL"   // School of Law
	CPS  College = "CPS"  // College of Professional Studies
	CS   College = "CS"   // College of Science
	CSSH College = "CSSH" // College of Social Sciences and Humanities
)

type Year uint

const (
	First    Year = 1
	Second   Year = 2
	Third    Year = 3
	Fourth   Year = 4
	Fifth    Year = 5
	Graduate Year = 6
)

type User struct {
	Model

	Role         UserRole `gorm:"type:varchar(255);" json:"user_role,omitempty" validate:"required,max=255"`
	NUID         string   `gorm:"column:nuid;type:varchar(9);unique" json:"nuid" validate:"required,numeric,len=9"`
	FirstName    string   `gorm:"type:varchar(255)" json:"first_name" validate:"required,max=255"`
	LastName     string   `gorm:"type:varchar(255)" json:"last_name" validate:"required,max=255"`
	Email        string   `gorm:"type:varchar(255);unique" json:"email" validate:"required,email,max=255"`
	PasswordHash string   `gorm:"type:varchar(97)" json:"-" validate:"required,len=97"`
	College      College  `gorm:"type:varchar(255)" json:"college" validate:"required,max=255"`
	Year         Year     `gorm:"type:smallint" json:"year" validate:"required,min=1,max=6"`

	Tag               []Tag     `gorm:"many2many:user_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Member            []Club    `gorm:"many2many:user_club_members;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Follower          []Club    `gorm:"many2many:user_club_followers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	IntendedApplicant []Club    `gorm:"many2many:user_club_intended_applicants;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Asked             []Comment `gorm:"foreignKey:AskedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-" validate:"-"`
	Answered          []Comment `gorm:"foreignKey:AnsweredByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-" validate:"-"`
	RSVP              []Event   `gorm:"many2many:user_event_rsvps;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Waitlist          []Event   `gorm:"many2many:user_event_waitlists;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type CreateUserRequestBody struct {
	NUID      string  `json:"nuid" validate:"required,numeric,len=9"`
	FirstName string  `json:"first_name" validate:"required,max=255"`
	LastName  string  `json:"last_name" validate:"required,max=255"`
	Email     string  `json:"email" validate:"required,email,neu_email,max=255"`
	Password  string  `json:"password" validate:"required,password"`
	College   College `json:"college" validate:"required,oneof=CAMD DMSB KCCS CE BCHS SL CPS CS CSSH"`
	Year      Year    `json:"year" validate:"required,min=1,max=6"`
}

type UpdateUserRequestBody struct {
	NUID      string  `json:"nuid" validate:"omitempty,numeric,len=9"`
	FirstName string  `json:"first_name" validate:"omitempty,max=255"`
	LastName  string  `json:"last_name" validate:"omitempty,max=255"`
	Email     string  `json:"email" validate:"omitempty,email,neu_email,max=255"`
	Password  string  `json:"password" validate:"omitempty,password"`
	College   College `json:"college" validate:"omitempty,oneof=CAMD DMSB KCCS CE BCHS SL CPS CS CSSH"`
	Year      Year    `json:"year" validate:"omitempty,min=1,max=6"`
}

type CreateUserTagsBody struct {
	Tags      []uuid.UUID  `json:"tags" validate:"required"`
}
