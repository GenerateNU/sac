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
	App      *fiber.App
	Address  string
	Conn     *gorm.DB
	Settings config.Settings
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

	configuration.Database.DatabaseName = generateRandomDBName()

	connectionWithDB, err := configureDatabase(configuration)

	if err != nil {
		return TestApp{}, err
	}

	return TestApp{
		App:     server.Init(connectionWithDB),
		Address: fmt.Sprintf("http://%s", listener.Addr().String()),
		Conn:    connectionWithDB,
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
	app.Conn.Exec(fmt.Sprintf("DROP DATABASE %s;", app.Settings.Database.DatabaseName))
}

type TestRequest struct {
	TestApp TestApp
	Assert  *assert.A
	Resp    *http.Response
}

func RequestTestSetup(t *testing.T, method string, path string, body *map[string]interface{}, headers *map[string]string) (TestApp, *assert.A, *http.Response) {
	app, assert := InitTest(t)

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
