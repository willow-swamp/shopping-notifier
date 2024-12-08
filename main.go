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
	http.HandleFunc("/items", item_handler.GetItems)
	http.HandleFunc("/item", item_handler.GetItem)
	//http.HandleFunc("/new", New)
	http.HandleFunc("/edit", item_handler.EditItem)
	http.HandleFunc("/create_item", item_handler.CreateItem)
	http.HandleFunc("/update_item", item_handler.UpdateItem)
	http.HandleFunc("/delete_item", item_handler.DeleteItem)

	user_service := service.NewUserService(dbRepository)
	user_handler := handler.NewUserHandler(user_service)

	http.HandleFunc("/users", user_handler.GetUsers)
	http.HandleFunc("/user", user_handler.GetUser)
	http.HandleFunc("/create_user", user_handler.CreateUser)

	http.ListenAndServe(":8080", nil)
}
