package model

import (
	"lottery/model/repository"
	"github.com/jinzhu/gorm"
	"time"
)

// 賞品情報
type Prize struct {
	repository.PgModel

	Quantity  uint64 `gorm:"not null"`
	Schedules []PrizeSchedule
}

func MigratePrize() {
	repository.Invoke(&Prize{}, func(db *gorm.DB) {
		db.AutoMigrate(&Prize{})
	})
}

func NewPrize() *Prize {
	return &Prize{}
}

func (p *Prize) Find(PrizeID uint64) *Prize {
	var Prizes []Prize
	repository.Invoke(&Prize{}, func(db *gorm.DB) {
		tim := time.Now()
		db.Preload("Schedules", "start_time <= ? and expire_time > ?", tim, tim).Where("id", PrizeID).Find(&Prizes)
	})
	if len(Prizes) <= 0 {
		return nil
	}
	return &Prizes[0]
}
