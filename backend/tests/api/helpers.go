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

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/database"
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/server"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitTest(t *testing.T) (TestApp, *assert.A) {
	assert := assert.New(t)
	app, err := spawnApp()

	assert.NilError(err)

	return app, assert
}

type TestApp struct {
	App           *fiber.App
	Address       string
	Conn          *gorm.DB
	Settings      config.Settings
	InitialDBName string
}

func spawnApp() (TestApp, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		return TestApp{}, err
	}

	configuration, err := config.GetConfiguration(filepath.Join("..", "..", "..", "config"))

	if err != nil {
		return TestApp{}, err
	}

	initialDBName := configuration.Database.DatabaseName

	configuration.Database.DatabaseName = generateRandomDBName()

	connectionWithDB, err := configureDatabase(configuration)

	if err != nil {
		return TestApp{}, err
	}

	return TestApp{
		App:           server.Init(connectionWithDB),
		Address:       fmt.Sprintf("http://%s", listener.Addr().String()),
		Conn:          connectionWithDB,
		Settings:      configuration,
		InitialDBName: initialDBName,
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
	Method  string
	Path    string
	Body    *map[string]interface{}
	Headers *map[string]string
}

func (request TestRequest) Test(t *testing.T, existingAppAssert *ExistingAppAssert) (ExistingAppAssert, *http.Response) {
	var app TestApp
	var assert *assert.A

	if existingAppAssert == nil {
		app, assert = InitTest(t)
	} else {
		app, assert = existingAppAssert.App, existingAppAssert.Assert
	}

	address := fmt.Sprintf("%s%s", app.Address, request.Path)

	var req *http.Request

	if request.Body == nil {
		t.Log("NO REQUEST BODY!")
		req = httptest.NewRequest(request.Method, address, nil)
	} else {
		bodyBytes, err := json.Marshal(request.Body)

		assert.NilError(err)

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

	resp, err := app.App.Test(req)

	assert.NilError(err)

	return ExistingAppAssert{
		App:    app,
		Assert: assert,
	}, resp
}

func (request TestRequest) TestOnStatus(t *testing.T, existingAppAssert *ExistingAppAssert, status int) ExistingAppAssert {
	appAssert, resp := request.Test(t, existingAppAssert)

	_, assert := appAssert.App, appAssert.Assert

	assert.Equal(status, resp.StatusCode)

	return appAssert
}

func (request TestRequest) TestOnError(t *testing.T, existingAppAssert *ExistingAppAssert, expectedError errors.Error) ExistingAppAssert {
	appAssert, resp := request.Test(t, existingAppAssert)
	assert := appAssert.Assert

	var respBody map[string]interface{}

	err := json.NewDecoder(resp.Body).Decode(&respBody)

	assert.NilError(err)
	assert.Equal(expectedError.Message, respBody["error"].(string))

	assert.Equal(expectedError.StatusCode, resp.StatusCode)

	return appAssert
}

type ErrorWithDBTester struct {
	Error    errors.Error
	DBTester DBTester
}

func (request TestRequest) TestOnErrorAndDB(t *testing.T, existingAppAssert *ExistingAppAssert, errorWithDBTester ErrorWithDBTester) ExistingAppAssert {
	appAssert := request.TestOnError(t, existingAppAssert, errorWithDBTester.Error)
	errorWithDBTester.DBTester(appAssert.App, appAssert.Assert, nil)
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
