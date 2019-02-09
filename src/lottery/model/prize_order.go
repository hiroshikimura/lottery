package model

import (
	"lottery/model/repository"
	"github.com/jinzhu/gorm"
	. "github.com/mattn/go-try/try"
)

// 抽選申し込み
type PrizeOrder struct {
	repository.PgModel

	// 抽選申し込みUUID
	OrderUUID string `gorm:"type:varchar(255);unique_index;not null"`

	// 抽選対象
	PrizeScheduleID uint64 `gorm:"not null"`

	// 抽選申し込み結果
	Result PrizeOrderResult
}

func MigratePrizeOrder() {
	repository.Invoke(&PrizeOrder{}, func(db *gorm.DB) {
		db.AutoMigrate(&PrizeOrder{})
	})
}

func (p *PrizeOrder) Add() error {
	var err error
	repository.Invoke(p, func(db *gorm.DB) {
		tx := db.Begin()
		Try(func() {
			db.Create(p)
			tx.Commit()
			db.Where("OrderUUID = ?", p.OrderUUID).Find(p)
		}).Catch(func(e error) {
			tx.Rollback()
			err = e
		})
	})
	return err
}
