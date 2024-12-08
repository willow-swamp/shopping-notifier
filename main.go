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
	service := service.NewService(dbRepository)
	handler := handler.NewHandler(service)

	// Start the server
	http.HandleFunc("/", handler.GetItems)
	http.HandleFunc("/show", handler.GetItem)
	//http.HandleFunc("/new", New)
	http.HandleFunc("/edit", handler.EditItem)
	http.HandleFunc("/create", handler.CreateItem)
	http.HandleFunc("/update", handler.UpdateItem)
	http.HandleFunc("/delete", handler.DeleteItem)
	http.ListenAndServe(":8080", nil)
}
