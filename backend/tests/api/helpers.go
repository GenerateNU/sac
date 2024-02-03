package tests

import (
	"bytes"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GenerateNU/sac/backend/src/auth"
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/database"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/GenerateNU/sac/backend/src/server"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthLevel string

var (
	SuperUser   AuthLevel = "super_user"
	StudentUser AuthLevel = "sample_user"
	LoggedOut   AuthLevel = "logged_out"
)

type TestUser struct {
	Email        string
	Password     string
	AccessToken  string
	RefreshToken string
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

func (app *TestApp) Auth(level AuthLevel) {
	if level == SuperUser {
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
			Email:        email,
			Password:     password,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
	} else if level == StudentUser {
		studentUser, rawPassword := SampleStudentFactory()

		_, err := app.Send(TestRequest{
			Method: fiber.MethodPost,
			Path:   "/api/v1/users/",
			Body:   SampleStudentJSONFactory(studentUser, rawPassword),
		})
		if err != nil {
			panic(err)
		}

		resp, err := app.Send(TestRequest{
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
			Email:        studentUser.Email,
			Password:     rawPassword,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
	}
}

func InitTest(t *testing.T) (TestApp, *assert.A) {
	assert := assert.New(t)
	app, err := spawnApp()

	assert.NilError(err)

	return *app, assert
}

type TestApp struct {
	App      *fiber.App
	Address  string
	Conn     *gorm.DB
	Settings config.Settings
	TestUser *TestUser
}

func spawnApp() (*TestApp, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	configuration, err := config.GetConfiguration(filepath.Join("..", "..", "..", "config"))
	if err != nil {
		return nil, err
	}

	configuration.Database.DatabaseName = generateRandomDBName()

	connectionWithDB, err := configureDatabase(configuration)
	if err != nil {
		return nil, err
	}

	return &TestApp{
		App:      server.Init(connectionWithDB, configuration),
		Address:  fmt.Sprintf("http://%s", listener.Addr().String()),
		Conn:     connectionWithDB,
		Settings: configuration,
	}, nil
}

func generateRandomInt(max int64) int64 {
	randInt, _ := crand.Int(crand.Reader, big.NewInt(max))
	return randInt.Int64()
}

func generateRandomDBName() string {
	prefix := "sac_test_"
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	length := len(prefix) + 36
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letterBytes[generateRandomInt(int64(len(letterBytes)))]
	}

	return fmt.Sprintf("%s%s", prefix, string(result))
}

func configureDatabase(config config.Settings) (*gorm.DB, error) {
	dsnWithoutDB := config.Database.WithoutDb()
	dbWithoutDB, err := gorm.Open(gormPostgres.Open(dsnWithoutDB), &gorm.Config{SkipDefaultTransaction: true, TranslateError: true})
	if err != nil {
		return nil, err
	}

	err = dbWithoutDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", config.Database.DatabaseName)).Error
	if err != nil {
		return nil, err
	}

	dsnWithDB := config.Database.WithDb()
	dbWithDB, err := gorm.Open(gormPostgres.Open(dsnWithDB), &gorm.Config{SkipDefaultTransaction: true, TranslateError: true})
	if err != nil {
		return nil, err
	}

	err = dbWithDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		return nil, err
	}
	err = database.MigrateDB(config, dbWithDB)
	if err != nil {
		return nil, err
	}

	return dbWithDB, nil
}

type ExistingAppAssert struct {
	App    TestApp
	Assert *assert.A
}

func (eaa ExistingAppAssert) Close() {
	db, err := eaa.App.Conn.DB()
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}

type TestRequest struct {
	Method    string
	Path      string
	Body      *map[string]interface{}
	Headers   *map[string]string
	AuthLevel *AuthLevel
}

func (app TestApp) Send(request TestRequest) (*http.Response, error) {
	address := fmt.Sprintf("%s%s", app.Address, request.Path)

	var req *http.Request

	if request.Body == nil {
		req = httptest.NewRequest(request.Method, address, nil)
	} else {
		bodyBytes, err := json.Marshal(request.Body)
		if err != nil {
			return nil, err
		}

		req = httptest.NewRequest(request.Method, address, bytes.NewBuffer(bodyBytes))

		if request.Headers == nil {
			request.Headers = &map[string]string{}
		}

		if _, ok := (*request.Headers)["Content-Type"]; !ok {
			(*request.Headers)["Content-Type"] = "application/json"
		}
	}

	if request.Headers != nil {
		for key, value := range *request.Headers {
			req.Header.Add(key, value)
		}
	}

	if app.TestUser != nil {
		req.AddCookie(&http.Cookie{
			Name:  "access_token",
			Value: app.TestUser.AccessToken,
		})
	}

	resp, err := app.App.Test(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (request TestRequest) Test(t *testing.T, existingAppAssert *ExistingAppAssert) (ExistingAppAssert, *http.Response) {
	if existingAppAssert == nil {
		app, assert := InitTest(t)

		if request.AuthLevel != nil {
			app.Auth(*request.AuthLevel)
		}
		existingAppAssert = &ExistingAppAssert{
			App:    app,
			Assert: assert,
		}
	}

	resp, err := existingAppAssert.App.Send(request)

	existingAppAssert.Assert.NilError(err)

	return *existingAppAssert, resp
}

func (request TestRequest) TestOnStatus(t *testing.T, existingAppAssert *ExistingAppAssert, status int) ExistingAppAssert {
	appAssert, resp := request.Test(t, existingAppAssert)

	_, assert := appAssert.App, appAssert.Assert

	assert.Equal(status, resp.StatusCode)

	return appAssert
}

func (request *TestRequest) testOn(t *testing.T, existingAppAssert *ExistingAppAssert, status int, key string, value string) (ExistingAppAssert, *http.Response) {
	appAssert, resp := request.Test(t, existingAppAssert)
	assert := appAssert.Assert

	var respBody map[string]interface{}

	err := json.NewDecoder(resp.Body).Decode(&respBody)

	assert.NilError(err)
	assert.Equal(value, respBody[key].(string))

	assert.Equal(status, resp.StatusCode)
	return appAssert, resp
}

func (request TestRequest) TestOnError(t *testing.T, existingAppAssert *ExistingAppAssert, expectedError errors.Error) ExistingAppAssert {
	appAssert, _ := request.testOn(t, existingAppAssert, expectedError.StatusCode, "error", expectedError.Message)
	return appAssert
}

type ErrorWithDBTester struct {
	Error    errors.Error
	DBTester DBTester
}

func (request TestRequest) TestOnErrorAndDB(t *testing.T, existingAppAssert *ExistingAppAssert, errorWithDBTester ErrorWithDBTester) ExistingAppAssert {
	appAssert, resp := request.testOn(t, existingAppAssert, errorWithDBTester.Error.StatusCode, "error", errorWithDBTester.Error.Message)
	errorWithDBTester.DBTester(appAssert.App, appAssert.Assert, resp)
	return appAssert
}

func (request TestRequest) TestOnMessage(t *testing.T, existingAppAssert *ExistingAppAssert, status int, message string) ExistingAppAssert {
	request.testOn(t, existingAppAssert, status, "message", message)
	return *existingAppAssert
}

func (request TestRequest) TestOnMessageAndDB(t *testing.T, existingAppAssert *ExistingAppAssert, status int, message string, dbTester DBTester) ExistingAppAssert {
	appAssert, resp := request.testOn(t, existingAppAssert, status, "message", message)
	dbTester(appAssert.App, appAssert.Assert, resp)
	return appAssert
}

type DBTester func(app TestApp, assert *assert.A, resp *http.Response)

type DBTesterWithStatus struct {
	Status int
	DBTester
}

func (request TestRequest) TestOnStatusAndDB(t *testing.T, existingAppAssert *ExistingAppAssert, dbTesterStatus DBTesterWithStatus) ExistingAppAssert {
	appAssert, resp := request.Test(t, existingAppAssert)
	app, assert := appAssert.App, appAssert.Assert

	assert.Equal(dbTesterStatus.Status, resp.StatusCode)

	dbTesterStatus.DBTester(app, assert, resp)

	return appAssert
}

func generateCasingPermutations(word string, currentPermutation string, index int, results *[]string) {
	if index == len(word) {
		*results = append(*results, currentPermutation)
		return
	}

	generateCasingPermutations(word, currentPermutation+strings.ToLower(string(word[index])), index+1, results)
	generateCasingPermutations(word, currentPermutation+strings.ToUpper(string(word[index])), index+1, results)
}

func AllCasingPermutations(word string) []string {
	results := make([]string, 0)
	generateCasingPermutations(word, "", 0, &results)
	return results
}
