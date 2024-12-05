package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/willow-swamp/shopping-notifier/config"
	"github.com/willow-swamp/shopping-notifier/databases"
	"github.com/willow-swamp/shopping-notifier/models"
)

func IndexItems(w http.ResponseWriter, r *http.Request) {
	items, err := databases.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(items) == 0 {
		http.Error(w, "No items", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func ShowItem(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(nId)
	item, err := databases.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var item models.Item
		item.UserID = r.FormValue("user_id")
		item.Name = r.FormValue("name")
		item.Priority = convPriority(r.FormValue("priority"))
		item.StockStatus = convStockStatus(r.FormValue("stock_status"))
		fmt.Println(item)
		err := databases.CreateItem(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(nId)
	item, err := databases.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nId := r.FormValue("id")
		id, _ := strconv.Atoi(nId)
		item, err := databases.GetItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		item.UserID = r.FormValue("user_id")
		item.Name = r.FormValue("name")
		item.Priority = convPriority(r.FormValue("priority"))
		item.StockStatus = convStockStatus(r.FormValue("stock_status"))
		err = databases.UpdateItem(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nId := r.FormValue("id")
		id, _ := strconv.Atoi(nId)
		err := databases.DeleteItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func convPriority(p string) config.PriorityType {
	switch p {
	case "1":
		return config.PriorityLow
	case "3":
		return config.PriorityHigh
	default:
		return config.PriorityMedium
	}
}

func convStockStatus(s string) config.StockStatusType {
	switch s {
	case "1":
		return config.StockStatusInStock
	case "2":
		return config.StockStatusOutOfStock
	default:
		return config.StockStatusOutOfStock
	}
}
