package helpers

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	db, err := dbWithDB.DB()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", filepath.Join("..", "..", "migrations")),
		"postgres", driver)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil {
		return nil, err
	}

	err = database.CreateSuperUserIfNotExists(settings, dbWithDB)
	if err != nil {
		return nil, err
	}

	return dbWithDB, nil
}
