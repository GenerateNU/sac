package models

type WelcomeTask struct {
	Email    string `gorm:"primaryKey"`
	Name     string
	Attempts int
}
