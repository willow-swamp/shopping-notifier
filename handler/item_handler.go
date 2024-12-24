package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/willow-swamp/shopping-notifier/config"
	"github.com/willow-swamp/shopping-notifier/models"
	"github.com/willow-swamp/shopping-notifier/service"
)

type ItemHandler struct {
	item_service  *service.ItemService
	user_service  *service.UserService
	group_service *service.GroupService
	store         *sessions.CookieStore
}

func NewItemHandler(item_service *service.ItemService, user_service *service.UserService, group_service *service.GroupService, store *sessions.CookieStore) *ItemHandler {
	return &ItemHandler{item_service: item_service, user_service: user_service, group_service: group_service, store: store}
}

var tmpl = template.Must(template.ParseGlob("front/*.tmpl"))

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	session, err := h.store.Get(r, "line-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sub, ok := session.Values["sub"].(string)
	if !ok || sub == "" {
		var emptyItems []models.Item
		tmpl.ExecuteTemplate(w, "Items", emptyItems)
		return
	}

	var LoginUser models.LoginUser
	LoginUser.Sub = sub
	LoginUser.Name = session.Values["name"].(string)
	LoginUser.Picture = session.Values["picture"].(string)

	items, err := h.item_service.GetItems(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//if len(items) == 0 {
	//	http.Error(w, "No items", http.StatusNotFound)
	//	return
	//}
	//json.NewEncoder(w).Encode(items)
	tmpl.ExecuteTemplate(w, "Items", items)
}

func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(nId)
	item, err := h.item_service.GetItem(id)
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
		err := h.item_service.CreateItem(&item)
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
	item, err := h.item_service.GetItem(id)
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
		item, err := h.item_service.GetItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		group_id := r.FormValue("group_id")
		item.GroupID, _ = strconv.Atoi(group_id)
		item.Name = r.FormValue("name")
		item.Priority = convPriority(r.FormValue("priority"))
		item.StockStatus = convStockStatus(r.FormValue("stock_status"))
		err = h.item_service.UpdateItem(item)
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
		err := h.item_service.DeleteItem(id)
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
