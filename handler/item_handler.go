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

	items, err := h.item_service.GetItems(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type DisplayItem struct {
		ID          uint
		Name        string
		Priority    string
		StockStatus string
		GroupID     int
	}

	var displayItems []DisplayItem
	for _, item := range items {
		displayItems = append(displayItems, DisplayItem{
			ID:          item.ID,
			Name:        item.Name,
			Priority:    convToPriority(item.Priority),       // 優先度を日本語に変換
			StockStatus: convToStockStatus(item.StockStatus), // 在庫状況を日本語に変換
			GroupID:     item.GroupID,
		})
	}
	//if len(items) == 0 {
	//	http.Error(w, "No items", http.StatusNotFound)
	//	return
	//}
	//json.NewEncoder(w).Encode(items)
	tmpl.ExecuteTemplate(w, "Items", displayItems)
}

//func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
//	nId := r.URL.Query().Get("id")
//	id, _ := strconv.Atoi(nId)
//	item, err := h.item_service.GetItem(id)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	if item == nil {
//		http.Error(w, "No item", http.StatusInternalServerError)
//		return
//	}
//	json.NewEncoder(w).Encode(item)
//}

func (h *ItemHandler) NewItem(w http.ResponseWriter, r *http.Request) {
	session, err := h.store.Get(r, "line-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sub, ok := session.Values["sub"].(string)
	if !ok || sub == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tmpl.ExecuteTemplate(w, "NewItem", nil)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session, err := h.store.Get(r, "line-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sub, ok := session.Values["sub"].(string)
	if !ok || sub == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.user_service.GetUser(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var item models.Item
	item.GroupID = user.GroupID
	item.Name = r.FormValue("name")
	item.Priority = convFromPriority(r.FormValue("priority"))
	item.StockStatus = convFromStockStatus(r.FormValue("stock_status"))
	err = h.item_service.CreateItem(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/items", http.StatusSeeOther)
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
		item.Priority = convFromPriority(r.FormValue("priority"))
		item.StockStatus = convFromStockStatus(r.FormValue("stock_status"))
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

func convToPriority(p int) string {
	switch p {
	case 1:
		return config.PriorityLow
	case 3:
		return config.PriorityHigh
	default:
		return config.PriorityMedium
	}
}

func convToStockStatus(s int) string {
	switch s {
	case 1:
		return config.StockStatusInStock
	default:
		return config.StockStatusOutOfStock
	}
}

func convFromPriority(p string) int {
	switch p {
	case config.PriorityLow:
		return 1
	case config.PriorityHigh:
		return 3
	default:
		return 2
	}
}

func convFromStockStatus(s string) int {
	switch s {
	case config.StockStatusInStock:
		return 1
	default:
		return 2
	}
}
