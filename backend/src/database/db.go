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

	if err := MigrateDB(settings, db); err != nil {
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

func MigrateDB(settings config.Settings, db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Category{},
		&models.Club{},
		&models.Contact{},
		&models.Event{},
		&models.Notification{},
		&models.PointOfContact{},
		&models.Tag{},
		&models.User{},
	)

	if err != nil {
		return err
	}

	tx := db.Begin()

	if err := tx.Error; err != nil {
		return err
	}

	superUser := models.User{
		Role:         models.Super,
		NUID:         "000000000",
		Email:        "generatesac@gmail.com",
		PasswordHash: settings.SuperUser.Password, // TODO: hash this
		FirstName:    "SAC",
		LastName:     "Super",
		College:      models.KCCS,
		Year:         models.First,
	}
	if err := tx.Create(&superUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	superClub := models.Club{
		Name:        "SAC",
		Preview:     "SAC",
		Description: "SAC",
	}
	if err := tx.Create(&superClub).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&superClub).Association("Member").Append(&superUser); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&superClub).Update("num_members", gorm.Expr("num_members + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
