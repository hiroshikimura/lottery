package usecase

import (
	"lottery/model"
	"lottery/model/repository"
	"errors"
	"github.com/jinzhu/gorm"
	. "github.com/mattn/go-try/try"
)

// 抽選ロジックそのもの
type Lottery struct {
	PrizeID   uint64
	OrderUUID string
}

// 抽選ステータス
type LotteryResultStatus int

const (
	Na LotteryResultStatus = iota
	Loose
	Consalotation
	Win
)

func (c LotteryResultStatus) String() string {
	switch c {
	case Loose:
		return "loose"
	case Consalotation:
		return "consalotation"
	case Win:
		return "win"
	}
	return ""
}

func (p *Lottery) reserve() (*model.PrizeOrder, error) {

	var prizeSchedule *model.PrizeSchedule
	var prizeOrder model.PrizeOrder

	// Prizeを取得する
	var prize *model.Prize
	prize = (model.NewPrize()).Find(p.PrizeID)
	if prize == nil {
		return nil, errors.New("not content")
	}
	if len(prize.Schedules) == 0 {
		return nil, errors.New("not scheduled")
	}
	prizeSchedule = &prize.Schedules[0]
	prizeOrder.OrderUUID = p.OrderUUID
	prizeOrder.PrizeScheduleID = prizeSchedule.ID
	err := prizeOrder.Add()
	if err != nil {
		return nil, err
	}
	return &prizeOrder, nil
}

func Reserve(PrizeID uint64, OrderUUID string) (*Lottery, error) {
	var usecase *Lottery
	if false {
		// 指定した賞品は存在しない
		return nil, errors.New("not a prize")
	}
	usecase = &Lottery{PrizeID: PrizeID, OrderUUID: OrderUUID}

	_, err := usecase.reserve()
	if err != nil {
		return nil, err
	}
	return usecase, nil
}

func (p *Lottery) Pick(PrizeID uint64, OrderUUID string) (LotteryResultStatus, error) {
	// 抽選予約していない
	if false {
		return Na, errors.New("no have permission")
	}

	// ここからクリティカルセクション
	var PrizeOrder model.PrizeOrder
	repository.Invoke(PrizeOrder, func(db *gorm.DB) {
		tx := db.Begin()
		Try(func() {
			// 抽選番号取得準備

			// 抽選番号,抽選結果取得

			// 賞品獲得
			if false {

			} else {

			}

			tx.Commit()
		}).Catch(func(e error) {
			tx.Rollback()
		})
	})

	// ここまでクリティカルセクション

	return Win, nil
}

func Order(PrizeID uint64, OrderUUID string) (LotteryResultStatus, error) {
	// これ、あまり使わないかも
	var lottery *Lottery
	var status LotteryResultStatus
	var err error
	lottery, err = Reserve(PrizeID, OrderUUID)
	if err != nil {
		return Na, err
	}
	status, err = lottery.Pick(PrizeID, OrderUUID)
	if err != nil {
		return Na, err
	}
	return status, nil
}
