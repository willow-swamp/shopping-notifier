package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/willow-swamp/shopping-notifier/databases"
	"github.com/willow-swamp/shopping-notifier/databases/repository"
	"github.com/willow-swamp/shopping-notifier/handler"
	"github.com/willow-swamp/shopping-notifier/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := databases.DBConn()
	if err != nil {
		panic(err)
	}

	err = databases.Migrate(db)
	if err != nil {
		panic(err)
	}

	var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	dbRepository := repository.NewMySQLRepository(db)

	user_service := service.NewUserService(dbRepository)
	group_service := service.NewGroupService(dbRepository)

	item_service := service.NewItemService(dbRepository)
	item_handler := handler.NewItemHandler(item_service, user_service, group_service, store)

	line_handler := handler.NewLineHandler(store)

	// Start the server
	http.HandleFunc("/", item_handler.GetItems)
	//http.HandleFunc("/item", item_handler.GetItem)
	http.HandleFunc("/new", item_handler.NewItem)
	http.HandleFunc("/edit", item_handler.EditItem)
	http.HandleFunc("/create_item", item_handler.CreateItem)
	http.HandleFunc("/update_item", item_handler.UpdateItem)
	http.HandleFunc("/delete_item", item_handler.DeleteItem)

	//user_handler := handler.NewUserHandler(user_service)

	//http.HandleFunc("/users", user_handler.GetUsers)
	//http.HandleFunc("/user", user_handler.GetUser)
	//http.HandleFunc("/create_user", user_handler.CreateUser)

	http.HandleFunc("/login", line_handler.LiffLoginUser)

	http.ListenAndServe(":8080", nil)
}
