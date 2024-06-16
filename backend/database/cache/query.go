package cache

import (
	go_json "github.com/goccy/go-json"

	"gorm.io/gorm"
)

type Query[T any] struct {
	Dest         T
	RowsAffected int64
}

func (q *Query[T]) Marshal() ([]byte, error) {
	return go_json.Marshal(q)
}

func (q *Query[T]) Unmarshal(bytes []byte) error {
	return go_json.Unmarshal(bytes, q)
}

func (q *Query[T]) copyTo(dst *Query[any]) error {
	bytes, err := q.Marshal()
	if err != nil {
		return err
	}

	return dst.Unmarshal(bytes)
}

func (q *Query[T]) replaceOn(db *gorm.DB) {
	SetPointedValue(db.Statement.Dest, q.Dest)
	SetPointedValue(&db.Statement.RowsAffected, &q.RowsAffected)
}
