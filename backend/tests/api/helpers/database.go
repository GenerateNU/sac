package helpers

import (
	"fmt"
	"sync"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/database"
	"gorm.io/gorm"
)

var (
	rootConn *gorm.DB
	once     sync.Once
)

func RootConn(dbSettings config.DatabaseSettings) {
	once.Do(func() {
		var err error
		rootConn, err = database.EstablishConn(dbSettings.WithoutDb())
		if err != nil {
			panic(err)
		}
	})
}

func configureDatabase(settings config.Settings) (*gorm.DB, error) {
	RootConn(settings.Database)

	err := rootConn.Exec(fmt.Sprintf("CREATE DATABASE %s", settings.Database.DatabaseName)).Error
	if err != nil {
		return nil, err
	}

	dbWithDB, err := database.EstablishConn(settings.Database.WithDb())
	if err != nil {
		return nil, err
	}

	err = database.MigrateDB(settings, dbWithDB)
	if err != nil {
		return nil, err
	}

	return dbWithDB, nil
}
