package database

import "gorm.io/gorm"

type QueryConfigurator func(*gorm.DB) *gorm.DB

func PreloadFollwer() QueryConfigurator {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Follower")
	}
}

func PreloadMember() QueryConfigurator {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Member")
	}
}

func PreloadTag() QueryConfigurator {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Tag")
	}
}

func PreloadClub() QueryConfigurator {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Club")
	}
}
