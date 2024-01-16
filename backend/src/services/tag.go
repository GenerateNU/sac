package services

import (
	"backend/src/models"
	"backend/src/transactions"

	"gorm.io/gorm"
)

type TagServiceInterface interface {
	CreateTag(tag models.Tag) (models.Tag, error)
}

type TagService struct {
	DB *gorm.DB
}

func (t *TagService) CreateTag(tag models.Tag) (models.Tag, error) {
	return transactions.CreateTag(t.DB, tag)
}
