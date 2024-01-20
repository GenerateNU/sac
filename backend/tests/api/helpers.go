package tests

import (
	"backend/src/config"
	"backend/src/database"
	"backend/src/server"
	"bytes"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

	configuration, err := config.GetConfiguration("../../../config")

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
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	length := 36
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letterBytes[generateRandomInt(int64(len(letterBytes)))]
	}

	return string(result)
}

func configureDatabase(config config.Settings) (*gorm.DB, error) {
	dsnWithoutDB := config.Database.WithoutDb()
	dbWithoutDB, err := gorm.Open(gormPostgres.Open(dsnWithoutDB), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}

	err = dbWithoutDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", config.Database.DatabaseName)).Error
	if err != nil {
		return nil, err
	}

	dsnWithDB := config.Database.WithDb()
	dbWithDB, err := gorm.Open(gormPostgres.Open(dsnWithDB), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		return nil, err
	}

	err = database.MigrateDB(config, dbWithDB)

	if err != nil {
		return nil, err
	}

	return dbWithDB, nil
}

func (app TestApp) DropDB() {
	db, err := app.Conn.DB()

	if err != nil {
		panic(err)
	}

	db.Close()

	app.Conn, err = gorm.Open(gormPostgres.Open(app.Settings.Database.WithoutDb()), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		panic(err)
	}

	app.Conn.Exec(fmt.Sprintf("DROP DATABASE %s;", app.Settings.Database.DatabaseName))
}

type TestRequest struct {
	TestApp TestApp
	Assert  *assert.A
	Resp    *http.Response
}

func RequestTester(t *testing.T, method string, path string, body *map[string]interface{}, headers *map[string]string, exisitingApp *TestApp, exisitingAssert *assert.A) (TestApp, *assert.A, *http.Response) {
	var app TestApp
	var assert *assert.A

	if exisitingApp == nil || exisitingAssert == nil {
		app, assert = InitTest(t)
	} else {
		app, assert = *exisitingApp, exisitingAssert
	}

	address := fmt.Sprintf("%s%s", app.Address, path)

	var req *http.Request

	if body == nil {
		req = httptest.NewRequest(method, address, nil)
	} else {
		bodyBytes, err := json.Marshal(body)

		assert.NilError(err)

		req = httptest.NewRequest(method, address, bytes.NewBuffer(bodyBytes))

		if headers != nil {
			for key, value := range *headers {
				req.Header.Set(key, value)
			}
		}
	}

	resp, err := app.App.Test(req)

	assert.NilError(err)

	return app, assert, resp
}

func RequestTesterWithJSONBody(t *testing.T, method string, path string, body *map[string]interface{}, headers *map[string]string, exisitingApp *TestApp, exisitingAssert *assert.A) (TestApp, *assert.A, *http.Response) {
	if headers == nil {
		headers = &map[string]string{"Content-Type": "application/json"}
	} else if _, ok := (*headers)["Content-Type"]; !ok {
		(*headers)["Content-Type"] = "application/json"
	}

	return RequestTester(t, method, path, body, headers, exisitingApp, exisitingAssert)
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
