package helpers

import (
	"fmt"
	"net"
	"path/filepath"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/server"
	"github.com/gofiber/fiber/v2"
	"github.com/huandu/go-assert"
	"gorm.io/gorm"
)

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

	configuration, err := config.GetConfiguration(filepath.Join("..", "..", "..", "config"), false)
	if err != nil {
		return nil, err
	}

	configuration.Database.DatabaseName = generateRandomDBName()

	connectionWithDB, err := configureDatabase(*configuration)
	if err != nil {
		return nil, err
	}

	return &TestApp{
		App:      server.Init(connectionWithDB, *configuration),
		Address:  fmt.Sprintf("http://%s", listener.Addr().String()),
		Conn:     connectionWithDB,
		Settings: *configuration,
	}, nil
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
