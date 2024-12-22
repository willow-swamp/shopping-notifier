package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/willow-swamp/shopping-notifier/models"
	"github.com/willow-swamp/shopping-notifier/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		http.Error(w, "No items", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(nId)
	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "No item", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user models.User
		user.LineID = r.FormValue("line_id")
		nGroupId := r.FormValue("group_id")
		user.GroupID, _ = strconv.Atoi(nGroupId)
		err := h.service.CreateUser(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
