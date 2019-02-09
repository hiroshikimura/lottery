package infrastructure

import (
	"fmt"
)

type Test struct {
}

func (b *Test) Help() string {
	return "Database Migratte execute"
}

func (b *Test) Run(args []string) int {
	config := LoadConfig()
	fmt.Printf("PATH=%s\n", config.Path)
	return 0
}

func (b *Test) Synopsis() string {
	return "Database Migrating"
}
