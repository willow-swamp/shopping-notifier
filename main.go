package main

import (
	"net/http"

	"github.com/willow-swamp/shopping-notifier/databases"
	"github.com/willow-swamp/shopping-notifier/databases/repository"
	"github.com/willow-swamp/shopping-notifier/handler"
	"github.com/willow-swamp/shopping-notifier/service"
)

func main() {
	db, err := databases.DBConn()
	if err != nil {
		panic(err)
	}

	err = databases.Migrate(db)
	if err != nil {
		panic(err)
	}

	dbRepository := repository.NewMySQLRepository(db)
	item_service := service.NewItemService(dbRepository)
	item_handler := handler.NewItemHandler(item_service)

	// Start the server
	http.HandleFunc("/", item_handler.GetItems)
	http.HandleFunc("/show", item_handler.GetItem)
	//http.HandleFunc("/new", New)
	http.HandleFunc("/edit", item_handler.EditItem)
	http.HandleFunc("/create", item_handler.CreateItem)
	http.HandleFunc("/update", item_handler.UpdateItem)
	http.HandleFunc("/delete", item_handler.DeleteItem)
	http.ListenAndServe(":8080", nil)
}
