package database

import (
	"errors"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/constants"
	"github.com/GenerateNU/sac/backend/database/cache"
	"github.com/GenerateNU/sac/backend/entities/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConfigureDB(settings config.Settings) (*gorm.DB, error) {
	db, err := EstablishConn(settings.Database.WithDb(), WithLoggerInfo())
	if err != nil {
		return nil, err
	}

	cache := cache.NewGormCache("db_cache", cache.NewRedisClient(settings.DBCache), cache.CacheConfig{
		TTL:    constants.DB_CACHE_TTL,
		Prefix: "cache:",
	})

	if err := db.Use(cache); err != nil {
		return nil, err
	}

	if err := CreateSuperUserIfNotExists(settings.SuperUser, db); err != nil {
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
		PrepareStmt:            true,
	}

	for _, opt := range opts {
		rootConfig = opt(rootConfig)
	}

	db, err := gorm.Open(postgres.Open(dsn), &rootConfig)
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

	sqlDB.SetMaxIdleConns(constants.MAX_IDLE_CONNECTIONS)
	sqlDB.SetMaxOpenConns(constants.MAX_OPEN_CONNECTIONS)

	return nil
}

func CreateSuperUserIfNotExists(settings config.SuperUserSettings, db *gorm.DB) error {
	var superUser models.User

	if err := db.Where("role = ?", models.Super).First(&superUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := createSuperUser(settings, db); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func createSuperUser(settings config.SuperUserSettings, db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	superUser, err := SuperUser(settings)
	if err != nil {
		tx.Rollback()
		return err
	}

	var user models.User

	if err := db.Where("id = ?", superUser.ID).First(&user).Error; err != nil {
		tx := db.Begin().Session(&gorm.Session{SkipHooks: true})

		if err := tx.Error; err != nil {
			return err
		}

		superClub := SuperClub()
		if err := tx.Create(&superClub).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Model(&superClub).Update("num_members", gorm.Expr("num_members + 1")).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Create(&superUser).Error; err != nil {
			tx.Rollback()
			return err
		}

		SuperUserUUID = superUser.ID

		membership := models.Membership{
			ClubID:         superClub.ID,
			UserID:         superUser.ID,
			MembershipType: models.MembershipTypeAdmin,
		}

		if err := tx.Create(&membership).Error; err != nil {
			tx.Rollback()
			return err
		}

		follower := models.Follower{
			ClubID: superClub.ID,
			UserID: superUser.ID,
		}

		if err := tx.Create(&follower).Error; err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit().Error
	}

	return tx.Commit().Error
}
