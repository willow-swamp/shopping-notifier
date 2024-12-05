package main

import (
	"net/http"

	"github.com/willow-swamp/shopping-notifier/controller"
	"github.com/willow-swamp/shopping-notifier/databases"
)

func main() {
	err := databases.Migrate()
	if err != nil {
		panic(err)
	}

	// Start the server
	http.HandleFunc("/", controller.IndexItems)
	http.HandleFunc("/show", controller.ShowItem)
	//http.HandleFunc("/new", New)
	//http.HandleFunc("/edit", Edit)
	http.HandleFunc("/create", controller.CreateItem)
	http.HandleFunc("/update", controller.UpdateItem)
	http.HandleFunc("/delete", controller.DeleteItem)
	http.ListenAndServe(":8080", nil)
}
