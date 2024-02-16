package models

type File struct {
	Model

	FileName  string `gorm:"type:varchar(255)" json:"file_name"`
	FileSize  int64  `gorm:"type:bigint;default:0" json:"file_size"`
	FileData  []byte
	ObjectKey string `gorm:"type:varchar(255);unique" json:"object_key"`
	Tags      []*Tag `gorm:"many2many:file_tags;" json:"tags"`

	ClubID uint `gorm:"foreignKey:ClubID;" json:"-" validate:"min=1"`
	UserID uint `gorm:"foreignKey:UserID;" json:"-" validate:"min=1"`
}
