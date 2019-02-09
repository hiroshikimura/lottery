package infrastructure

import (
	"fmt"
	"lottery/model"
)

type Migarte struct {
}

func migrate() {
	fmt.Println(">>>>>>>>>>>>> migrate start")
	model.MigratePrize()
	model.MigratePrizeSchedule()
	model.MigratePrizeResultSummary()
	model.MigratePrizeOrder()
	model.MigratePrizeOrderResult()
	fmt.Println(">>>>>>>>>>>>> migrate end")
}

func (b *Migarte) Help() string {
	return "Database Migratte execute"
}

func (b *Migarte) Run(args []string) int {
	migrate()
	return 0
}

func (b *Migarte) Synopsis() string {
	return "Database Migrating"
}
