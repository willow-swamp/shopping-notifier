package main

import (
	"fmt"

	"github.com/willow-swamp/shopping-notifier/databases"
)

func main() {
	err := databases.Migrate()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Migrate success")
}
