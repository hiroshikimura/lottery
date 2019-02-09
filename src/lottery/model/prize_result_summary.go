package model

import (
	"lottery/model/repository"
	"github.com/jinzhu/gorm"
)

// 抽選結果サマリー
type PrizeResultSummary struct {
	repository.PgModel

	PrizeScheduleID uint64 `gorm:"not null"`

	// 抽選申し込み数
	OrderCount uint64 `gorm:"not null"`

	// 抽選大当たり数
	WinCount uint64 `gorm:"not null"`

	// 抽選中当たり数
	ConsolationCount uint64 `gorm:"not null"`
}

func MigratePrizeResultSummary() {
	repository.Invoke(&PrizeResultSummary{}, func(db *gorm.DB) {
		db.AutoMigrate(&PrizeResultSummary{})
	})
}
