package transactions

import "gorm.io/gorm"

type OptionalQuery func(*gorm.DB) *gorm.DB

func PreloadFollwer() OptionalQuery {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Follower")
	}
}

func PreloadMember() OptionalQuery {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Member")
	}
}

func PreloadTag() OptionalQuery {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Tag")
	}
}
