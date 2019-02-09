package model

import (
	"lottery/model/repository"
	"github.com/jinzhu/gorm"
)

// 賞品要素 削除予定
type PrizeElement struct {
	repository.PgModel

	OrderUUID string `gorm:"type:varchar(255);unique_index"`
}

func MigratePrizeElement() {
	repository.Invoke(&PrizeElement{}, func(db *gorm.DB) {
		db.AutoMigrate(&PrizeElement{})
	})
}
