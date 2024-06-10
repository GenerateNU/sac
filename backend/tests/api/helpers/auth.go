package helpers

import (
	"fmt"

	"github.com/GenerateNU/sac/backend/database"
	"github.com/GenerateNU/sac/backend/entities/models"
	"github.com/gofiber/fiber/v2"
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
	superUser, superUserErr := database.SuperUser(app.Settings.SuperUser)
	if superUserErr != nil {
		panic(fmt.Sprintf("Failed to get super user: %v", superUserErr))
	}

	email := superUser.Email
	password := app.Settings.SuperUser.Password

	resp, err := app.Send(TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/auth/login",
		Body: &map[string]interface{}{
			"email":    email,
			"password": password.Expose(),
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to authenticate super user: %v", err))
	}

	var accessToken string
	var refreshToken string

	authHeader := resp.Header.Get("Authorization")
	if authHeader != "" {
		accessToken = authHeader[len("Bearer "):]
	}

	if accessToken == "" || refreshToken == "" {
		panic("Failed to authenticate super user")
	}

	app.TestUser = &TestUser{
		UUID:        database.SuperUserUUID,
		Email:       email,
		AccessToken: accessToken,
	}
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
