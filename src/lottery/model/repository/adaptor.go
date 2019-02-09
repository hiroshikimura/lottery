package repository

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Adaptor interface {
	Open() *gorm.DB
}

// ID(PK)をuint64、AUTO INCREMENTに修正したbase model
type AbstractModel struct {
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func Invoke(p Adaptor, proc func(db *gorm.DB)) {
	db := p.Open()
	defer func() {
		db.Close()
	}()

	proc(db)
}
