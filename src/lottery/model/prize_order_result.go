package model

import (
	"lottery/model/repository"
	"github.com/jinzhu/gorm"
	"time"
)

// 抽選申し込み結果
type PrizeOrderResult struct {
	repository.PgModel
	PrizeOrderID uint64

	// 抽選申し込み日時
	OrderedAt time.Time `sql:"type:timestamptz"`

	// 抽選結果
	LotteryResult uint `gorm:"not null"`

	// 抽選番号
	LotteryNumber uint64 `gorm:"not null"`

	// 抽選時総申し込み数
	OrderCount uint64 `gorm:"not null"`

	// 抽選時大当たり当選数
	WinCount uint64 `gorm:"not null"`

	// 抽選時中当たり当選数
	ConsolationCount uint64 `gorm:"not null"`
}

func MigratePrizeOrderResult() {
	repository.Invoke(&PrizeOrderResult{}, func(db *gorm.DB) {
		db.AutoMigrate(&PrizeOrderResult{})
	})
}
