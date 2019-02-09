package repository

import (
	"github.com/jinzhu/gorm"
)

type MySQLModel struct {
	AbstractModel
}

func (p MySQLModel) Open() *gorm.DB {
	db, err := gorm.Open("mysql", "connection parameters")
	if err != nil {
		panic(err)
	}
	return db
}
