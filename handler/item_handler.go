package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/willow-swamp/shopping-notifier/config"
	"github.com/willow-swamp/shopping-notifier/models"
	"github.com/willow-swamp/shopping-notifier/service"
)

type ItemHandler struct {
	service *service.ItemService
}

func NewItemHandler(s *service.ItemService) *ItemHandler {
	return &ItemHandler{service: s}
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetItems()
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

func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(nId)
	item, err := h.service.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if item == nil {
		http.Error(w, "No item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var item models.Item
		group_id := r.FormValue("group_id")
		item.GroupID, _ = strconv.Atoi(group_id)
		item.Name = r.FormValue("name")
		item.Priority = convPriority(r.FormValue("priority"))
		item.StockStatus = convStockStatus(r.FormValue("stock_status"))
		err := h.service.CreateItem(&item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *ItemHandler) EditItem(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(nId)
	item, err := h.service.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if item == nil {
		http.Error(w, "No item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nId := r.FormValue("id")
		id, _ := strconv.Atoi(nId)
		item, err := h.service.GetItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		group_id := r.FormValue("group_id")
		item.GroupID, _ = strconv.Atoi(group_id)
		item.Name = r.FormValue("name")
		item.Priority = convPriority(r.FormValue("priority"))
		item.StockStatus = convStockStatus(r.FormValue("stock_status"))
		err = h.service.UpdateItem(item)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nId := r.FormValue("id")
		id, _ := strconv.Atoi(nId)
		err := h.service.DeleteItem(id)
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
