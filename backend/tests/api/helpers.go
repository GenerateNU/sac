package tests

import (
	"backend/src/config"
	"backend/src/database"
	"backend/src/models"
	"backend/src/server"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"net"

	"github.com/gofiber/fiber/v2"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestApp struct {
	App     *fiber.App
	Address string
	Conn    *gorm.DB
}

func SpawnApp() (TestApp, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		return TestApp{}, err
	}

	configuration, err := config.GetConfiguration("../../../config")

	if err != nil {
		return TestApp{}, err
	}

	configuration.Database.DatabaseName = generateRandomDBName()

	connectionWithDB, err := configureDatabase(configuration.Database)

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

func configureDatabase(config config.DatabaseSettings) (*gorm.DB, error) {
	dsnWithoutDB := config.WithoutDb()
	dbWithoutDB, err := gorm.Open(gormPostgres.Open(dsnWithoutDB), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return nil, err
	}

	err = dbWithoutDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", config.DatabaseName)).Error
	if err != nil {
		return nil, err
	}

	dsnWithDB := config.WithDb()
	dbWithDB, err := gorm.Open(gormPostgres.Open(dsnWithDB), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		return nil, err
	}

	err = database.MigrateDB(dbWithDB)

	if err != nil {
		return nil, err
	}

	return dbWithDB, nil
}

func (app *TestApp) InsertSampleUser() (models.User, error) {
	user := models.User{
		Role:         models.Super,
		NUID:         "000000000",
		Email:        "generatesac@gmail.com",
		PasswordHash: "rust",
		FirstName:    "SAC",
		LastName:     "Super",
		College:      models.KCCS,
		Year:         models.First,
	}

	result := app.Conn.Create(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
