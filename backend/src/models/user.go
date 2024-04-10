package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

var (
	Super   UserRole = "super"
	Student UserRole = "student"
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

type GraduationCycle string

const (
	December GraduationCycle = "december"
	May      GraduationCycle = "may"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	Model

	Role            UserRole        `gorm:"type:varchar(255);default:'student'" json:"role" validate:"required,oneof=super student"`
	FirstName       string          `gorm:"type:varchar(255)" json:"first_name" validate:"required,max=255"`
	LastName        string          `gorm:"type:varchar(255)" json:"last_name" validate:"required,max=255"`
	Email           string          `gorm:"type:varchar(255);unique" json:"email" validate:"required,email,max=255"`
	PasswordHash    string          `gorm:"type:varchar(97)" json:"-" validate:"required,len=97"`
	College         College         `gorm:"type:varchar(255)" json:"college" validate:"required,max=255"`
	GraduationCycle GraduationCycle `gorm:"type:varchar(255)" json:"graduation_cycle" validate:"required,max=255,oneof=december may"`
	GraduationYear  int16           `gorm:"type:smallint" json:"graduation_year" validate:"required"`
	IsVerified      bool            `gorm:"type:boolean;default:false" json:"is_verified"`

	Tag               []Tag     `gorm:"many2many:user_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Admin             []Club    `gorm:"many2many:user_club_admins;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Member            []Club    `gorm:"many2many:user_club_members;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Follower          []Club    `gorm:"many2many:user_club_followers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	IntendedApplicant []Club    `gorm:"many2many:user_club_intended_applicants;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Asked             []Comment `gorm:"foreignKey:AskedByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-" validate:"-"`
	Answered          []Comment `gorm:"foreignKey:AnsweredByID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-" validate:"-"`
	RSVP              []Event   `gorm:"many2many:user_event_rsvps;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
	Waitlist          []Event   `gorm:"many2many:user_event_waitlists;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-" validate:"-"`
}

type CreateUserRequestBody struct {
	FirstName string `json:"first_name" validate:"required,max=255"`
	LastName  string `json:"last_name" validate:"required,max=255"`
	Email     string `json:"email" validate:"required,email,neu_email,max=255"`
	Password  string `json:"password" validate:"required,password,min=8,max=255"`
	// Optional fields
	College         College         `json:"college" validate:"omitempty,oneof=CAMD DMSB KCCS CE BCHS SL CPS CS CSSH"`
	GraduationCycle GraduationCycle `json:"graduation_cycle" validate:"omitempty,max=255,oneof=december may"`
	GraduationYear  int16           `json:"graduation_year" validate:"omitempty"`
}

type UpdateUserRequestBody struct {
	FirstName       string          `json:"first_name" validate:"omitempty,max=255"`
	LastName        string          `json:"last_name" validate:"omitempty,max=255"`
	College         College         `json:"college" validate:"omitempty,oneof=CAMD DMSB KCCS CE BCHS SL CPS CS CSSH"`
	GraduationCycle GraduationCycle `json:"graduation_cycle" validate:"omitempty,max=255,oneof=december may"`
	GraduationYear  int16           `json:"graduation_year" validate:"omitempty"`
}

type LoginUserResponseBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatePasswordRequestBody struct {
	OldPassword string `json:"old_password" validate:"required,password,min=8,max=255"`
	NewPassword string `json:"new_password" validate:"required,password,nefield=OldPassword,min=8,max=255"`
}

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type CreateUserTagsBody struct {
	Tags []uuid.UUID `json:"tags" validate:"required"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	sac := &Club{}
	if err := tx.Where("name = ?", "SAC").First(sac).Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Member").Append(sac); err != nil {
		return err
	}

	if err := tx.Model(u).Association("Follower").Append(sac); err != nil {
		return err
	}

	if err := tx.Model(&Club{}).Where("id = ?", sac.ID).Update("num_members", gorm.Expr("num_members + 1")).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	sac := &Club{}
	if err := tx.Where("name = ?", "SAC").First(sac).Error; err != nil {
		return err
	}

	if err := tx.Model(&Club{}).Where("id = ?", sac.ID).Update("num_members", gorm.Expr("num_members - 1")).Error; err != nil {
		return err
	}

	return nil
}
