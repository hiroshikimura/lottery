package model

import (
	"lottery/model/repository"
	"github.com/jinzhu/gorm"
	"time"
)

// 抽選スケジュール情報
type PrizeSchedule struct {
	repository.PgModel

	PrizeID uint64
	// 開始日時
	StartTime time.Time `gorm:"not null;type:timestamptz"`
	// 終了日時
	ExpireTime time.Time `gorm:"not null;type:timestamptz"`
	// 大当たり当選数
	WinQuantity uint64 `gorm:"not null"`
	// 中当たり割合
	ConsolationRatio uint `gorm:"not null"`
	// 予想申し込み数
	ProspectiveQuantity uint64 `gorm:"not null"`

	Summary PrizeResultSummary
}

func MigratePrizeSchedule() {
	repository.Invoke(&PrizeSchedule{}, func(db *gorm.DB) {
		db.AutoMigrate(&PrizeSchedule{})
	})
}
