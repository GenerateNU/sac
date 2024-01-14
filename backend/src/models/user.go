package models

import "backend/src/types"

type UserRole string

const (
	Super     UserRole = "super"
	ClubAdmin UserRole = "clubAdmin"
	Student   UserRole = "student"
)

type College string

const (
	CAMD  College = "CAMD"  // College of Arts, Media and Design
	DMSB  College = "DMSB"  // D'Amore-McKim School of Business
	KCCS  College = "KCCS"  // Khoury College of Computer Sciences
	CoE   College = "CoE"   // College of Engineering
	BCoHS College = "BCoHS" // Bouvé College of Health Sciences
	SoL   College = "SoL"   // School of Law
	CoPS  College = "CoPS"  // College of Professional Studies
	CoS   College = "CoS"   // College of Science
	CoSSH College = "CoSSH" // College of Social Sciences and Humanities
)

type User struct {
	types.Model
	Role         UserRole `gorm:"type:user_role;" json:"user_role" validate:"required"`
	NUID         string   `gorm:"type:varchar(9);unique" json:"nuid" validate:"required"`
	FirstName    string   `gorm:"type:varchar(255)" json:"first_name" validate:"required"`
	LastName     string   `gorm:"type:varchar(255)" json:"last_name" validate:"required"`
	Email        string   `gorm:"type:varchar(255);unique" json:"email" validate:"required,email"`
	PasswordHash string   `gorm:"type:text" json:"-" validate:"-"`
	DateOfBirth  string   `gorm:"type:date" json:"date_of_birth" validate:"required"`
	College      College  `gorm:"type:college" json:"college" validate:"required"`
}
