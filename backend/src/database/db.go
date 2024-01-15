package database

import (
	"backend/src/config"
	"backend/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConfigureDB(settings config.Settings) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(settings.Database.WithDb()), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	if err := MigrateDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

func ConnPooling(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return nil
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Category{},
		&models.Club{},
		&models.Contact{},
		&models.Event{},
		&models.Notification{},
		&models.PointOfContact{},
		&models.Tag{},
		&models.User{},
	)
}
