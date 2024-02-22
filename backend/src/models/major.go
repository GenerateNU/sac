package models

type Majors string

type Major struct {
	Model

	Major string `gorm:"type:varchar(255);unique" json:"major" validate:"required,max=255"`

	Users []User `gorm:"many2many:user_majors;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"users,omitempty" validate:"-"`
}
