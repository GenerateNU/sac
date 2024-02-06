package helpers

import (
	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/database"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TestUser struct {
	UUID         uuid.UUID
	Email        string
	Password     string
	AccessToken  string
	RefreshToken string
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
		panic(superUserErr)
	}

	email := superUser.Email
	password := app.Settings.SuperUser.Password

	resp, err := app.Send(TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/auth/login",
		Body: &map[string]interface{}{
			"email":    email,
			"password": password,
		},
	})
	if err != nil {
		panic(err)
	}

	var accessToken string
	var refreshToken string

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "access_token" {
			accessToken = cookie.Value
		} else if cookie.Name == "refresh_token" {
			refreshToken = cookie.Value
		}
	}

	if accessToken == "" || refreshToken == "" {
		panic("Failed to authenticate super user")
	}

	app.TestUser = &TestUser{
		UUID:         database.SuperUserUUID,
		Email:        email,
		Password:     password.Expose(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func (app *TestApp) authStudent() {
	studentUser, rawPassword := SampleStudentFactory()

	resp, err := app.Send(TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/users/",
		Body:   SampleStudentJSONFactory(studentUser, rawPassword),
	})
	if err != nil {
		panic(err)
	}
	var respBody map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		panic(err)
	}

	rawStudentUserUUID := respBody["id"].(string)
	studentUserUUID, err := uuid.Parse(rawStudentUserUUID)
	if err != nil {
		panic(err)
	}

	resp, err = app.Send(TestRequest{
		Method: fiber.MethodPost,
		Path:   "/api/v1/auth/login",
		Body: &map[string]interface{}{
			"email":    studentUser.Email,
			"password": rawPassword,
		},
	})
	if err != nil {
		panic(err)
	}

	var accessToken string
	var refreshToken string

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "access_token" {
			accessToken = cookie.Value
		} else if cookie.Name == "refresh_token" {
			refreshToken = cookie.Value
		}
	}

	if accessToken == "" || refreshToken == "" {
		panic("Failed to authenticate sample student user")
	}

	app.TestUser = &TestUser{
		UUID:         studentUserUUID,
		Email:        studentUser.Email,
		Password:     rawPassword,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func SampleStudentFactory() (models.User, string) {
	password := "1234567890&"
	hashedPassword, err := auth.ComputePasswordHash(password)
	if err != nil {
		panic(err)
	}

	return models.User{
		Role:         models.Student,
		FirstName:    "Jane",
		LastName:     "Doe",
		Email:        "doe.jane@northeastern.edu",
		PasswordHash: *hashedPassword,
		NUID:         "001234567",
		College:      models.KCCS,
		Year:         models.Third,
	}, password
}

func SampleStudentJSONFactory(sampleStudent models.User, rawPassword string) *map[string]interface{} {
	if sampleStudent.Role != models.Student {
		panic("User is not a student")
	}
	return &map[string]interface{}{
		"first_name": sampleStudent.FirstName,
		"last_name":  sampleStudent.LastName,
		"email":      sampleStudent.Email,
		"password":   rawPassword,
		"nuid":       sampleStudent.NUID,
		"college":    string(sampleStudent.College),
		"year":       int(sampleStudent.Year),
	}
}
