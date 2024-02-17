package database

import (
	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConfigureDB(settings config.Settings) (*gorm.DB, error) {
	db, err := EstablishConn(settings.Database.WithDb(), WithLoggerInfo())
	if err != nil {
		return nil, err
	}

	if err := MigrateDB(settings, db); err != nil {
		return nil, err
	}

	return db, nil
}

type OptionalFunc func(gorm.Config) gorm.Config

func WithLoggerInfo() OptionalFunc {
	return func(gormConfig gorm.Config) gorm.Config {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
		return gormConfig
	}
}

func EstablishConn(dsn string, opts ...OptionalFunc) (*gorm.DB, error) {
	rootConfig := gorm.Config{
		SkipDefaultTransaction: true,
		TranslateError:         true,
	}

	for _, opt := range opts {
		rootConfig = opt(rootConfig)
	}

	db, err := gorm.Open(postgres.Open(dsn), &rootConfig)
	if err != nil {
		return nil, err
	}

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
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
		&models.File{}, 
		&models.Series{},
		&models.EventInstanceException{},
		&models.EventSeries{},
		&models.Membership{},
	)
	if err != nil {
		return err
	}

	// Check if the database already has a super user
	var superUser models.User
	if err := db.Where("role = ?", models.Super).First(&superUser).Error; err != nil {
		if err := createSuperUser(settings, db); err != nil {
			return err
		}
	}

	return nil
}

func createSuperUser(settings config.Settings, db *gorm.DB) error {
	tx := db.Begin()

	if err := tx.Error; err != nil {
		return err
	}

	superUser, err := SuperUser(settings.SuperUser)
	if err != nil {
		tx.Rollback()
		return err
	}

	var user models.User

	if err := db.Where("nuid = ?", superUser.NUID).First(&user).Error; err != nil {
		tx := db.Begin()

		if err := tx.Error; err != nil {
			return err
		}

		if err := tx.Create(&superUser).Error; err != nil {
			tx.Rollback()
			return err
		}

		SuperUserUUID = superUser.ID

		superClub := SuperClub()

		if err := tx.Create(&superClub).Error; err != nil {
			tx.Rollback()
			return err
		}

		membership := models.Membership{
			ClubID:         superClub.ID,
			UserID:         superUser.ID,
			MembershipType: models.MembershipTypeAdmin,
		}

		if err := tx.Create(&membership).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	}
	return nil
}
