package models

import (
	"encoding/gob"

	go_json "github.com/goccy/go-json"
	"gorm.io/gorm"
)

func init() {
	gob.Register(User{})
}

type User struct {
	Model

	Role            UserRole        `json:"role"`
	Name            string          `json:"name"`
	Email           string          `json:"email"`
	Major0          Major           `json:"major0"`
	Major1          Major           `json:"major1"`
	Major2          Major           `json:"major2"`
	College         College         `json:"college"`
	GraduationCycle GraduationCycle `json:"graduation_cycle"`
	GraduationYear  int16           `json:"graduation_year"`
	AvatarURL       string          `json:"avatar_url"`

	Tag               []Tag     `gorm:"many2many:user_tags;" json:"-"`
	Admin             []Club    `gorm:"many2many:user_club_admins;" json:"-"`
	Member            []Club    `gorm:"many2many:user_club_members;" json:"-"`
	Follower          []Club    `gorm:"many2many:user_club_followers;" json:"-"`
	IntendedApplicant []Club    `gorm:"many2many:user_club_intended_applicants;" json:"-"`
	Asked             []Comment `gorm:"foreignKey:AskedByID;" json:"-"`
	Answered          []Comment `gorm:"foreignKey:AnsweredByID;" json:"-"`
	RSVP              []Event   `gorm:"many2many:user_event_rsvps;" json:"-"`
	Waitlist          []Event   `gorm:"many2many:user_event_waitlists;" json:"-"`
}

func (u *User) Marshal() string {
	user, _ := go_json.Marshal(u)
	return string(user)
}

func UnmarshalUser(user string) User {
	u := User{}
	go_json.Unmarshal([]byte(user), &u)
	return u
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
