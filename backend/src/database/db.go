package database

import (
	"backend/src/config"
	"backend/src/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreatePostgresConnection(settings config.Settings) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(settings.Database.WithDb()), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	if err := db.AutoMigrate(
		&models.User{},
	); err != nil {
		return nil, fmt.Errorf("failed to perform database auto migration: %v", err)
	}

	return db, nil
}
