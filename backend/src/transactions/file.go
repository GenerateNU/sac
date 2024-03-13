package transactions

import (
	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OwnerID   uuid.UUID `gorm:"uniqueIndex:compositeindex;index;not null;foreignKey:OwnerID" json:"-" validate:"uuid4"`
// OwnerType string    `gorm:"uniqueIndex:compositeindex;index;not null;type:varchar(255)" json:"-" validate:"required,max=255"`

// FileName string `gorm:"type:varchar(255)" json:"file_name" validate:"required,max=255"`
// FileType string `gorm:"type:varchar(255)" json:"file_type" validate:"required,max=255"`
// FileSize int    `gorm:"type:int" json:"file_size" validate:"required,min=1"`
// FileURL  string `gorm:"type:varchar(255)" json:"file_url" validate:"required,max=255"`
// ObjectKey string `gorm:"type:varchar(255)" json:"object_key" validate:"required,max=255"`

func CreateFile(db *gorm.DB, ownerID uuid.UUID, ownerType string, fileURL string) (*models.File, *errors.Error) {
	file := &models.File{
		OwnerID:   ownerID,
		OwnerType: ownerType,
		FileURL:   fileURL,
		FileType:  "fileType",
		FileSize:  0,
		FileName:  "fileName",
	}

	if err := db.Create(file).Error; err != nil {
		return nil, &errors.FailedToCreateFileInDB
	}

	return file, nil
}