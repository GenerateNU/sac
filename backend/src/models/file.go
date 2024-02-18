package models

import "github.com/google/uuid"

type File struct {
	Model

	FileName  string `gorm:"type:varchar(255)" json:"file_name"`
	FileSize  int64  `gorm:"type:bigint;default:0" json:"file_size"`
	FileData  []byte
	ObjectKey string `gorm:"type:varchar(255);unique" json:"object_key"`
	Tags      []*Tag `gorm:"many2many:file_tags;" json:"tags"`
	S3Url     string `gorm:"type:varchar(255);default:NULL" json:"photo" validate:"url,max=255"`

	AssociationType string    `gorm:"type:varchar(255)" json:"association_type"`                // association with files (club/event/user)
	AssociationID   uuid.UUID `gorm:"type:varchar(255)" json:"association_id" validate:"min=1"` // association id (club/event/user)
}

type FileBody struct {
	AssociationType string    `gorm:"type:varchar(255)" json:"association_type"`                // association with files (club/event/user)
	AssociationID   uuid.UUID `gorm:"type:varchar(255)" json:"association_id" validate:"min=1"` // association id (club/event/user)
}
