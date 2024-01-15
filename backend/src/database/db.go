package database

import (
	"backend/src/config"
	"backend/src/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreatePostgresConnection(settings config.Settings) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(settings.Database.WithDb()), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
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
		&models.Category{},
		&models.Club{},
		&models.Contact{},
		&models.Event{},
		&models.Notification{},
		&models.PointOfContact{},
		&models.Tag{},
		&models.User{},
	); err != nil {
		return nil, fmt.Errorf("failed to perform database auto migration: %v", err)
	}

	return db, nil
}
