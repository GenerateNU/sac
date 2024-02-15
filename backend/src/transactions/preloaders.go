package transactions

import "gorm.io/gorm"

type OptionalPreload func(*gorm.DB) *gorm.DB

func PreloadFollwer() OptionalPreload {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Follower")
	}
}

func PreloadMember() OptionalPreload {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Member")
	}
}
