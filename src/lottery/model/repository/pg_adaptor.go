package repository

import (
	"github.com/jinzhu/gorm"
)

type PgModel struct {
	AbstractModel
}

func (p PgModel) Open() *gorm.DB {
	db, err := gorm.Open("postgres", "host=dev.local port=5432 user=kimura dbname=kimura password=*****************")
	if err != nil {
		panic(err)
	}
	return db
}
