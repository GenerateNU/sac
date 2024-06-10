package helpers

import (
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/google/uuid"
)

type TestUser struct {
	UUID        uuid.UUID
	Email       string
	AccessToken string
}

func (app *TestApp) Auth(role models.UserRole) {
	if role == models.Super {
		app.authSuper()
	} else if role == models.Student {
		app.authStudent()
	}
	// unauthed -> do nothing
}

func (app *TestApp) authSuper() {
	app.TestUser = nil
}

func (app *TestApp) authStudent() {
	app.TestUser = nil
}

func SampleStudentFactory() models.User {
	return models.User{
		Role:            models.Student,
		Name:            "Jane Doe",
		Email:           "doe.jane@northeastern.edu",
		Major0:          models.ComputerScience,
		Major1:          models.Economics,
		College:         models.KCCS,
		GraduationCycle: models.May,
		GraduationYear:  2025,
	}
}

func SampleStudentJSONFactory(sampleStudent models.User) *map[string]interface{} {
	if sampleStudent.Role != models.Student {
		panic("User is not a student")
	}
	return &map[string]interface{}{
		"name":             sampleStudent.Name,
		"email":            sampleStudent.Email,
		"major0":           string(sampleStudent.Major0),
		"major1":           string(sampleStudent.Major1),
		"major2":           string(sampleStudent.Major2),
		"college":          string(sampleStudent.College),
		"graduation_cycle": string(sampleStudent.GraduationCycle),
		"graduation_year":  int(sampleStudent.GraduationYear),
	}
}
