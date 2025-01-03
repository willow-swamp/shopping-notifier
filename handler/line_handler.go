package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/willow-swamp/shopping-notifier/line"
)

type LineHandler struct {
	store *sessions.CookieStore
}

func NewLineHandler(store *sessions.CookieStore) *LineHandler {
	return &LineHandler{store: store}
}

func (lineHandler *LineHandler) LiffLoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		IDToken string `json:"id_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lineAccess := line.NewLineAccess(req.IDToken)
	userData := lineAccess.GetUserData()
	if userData == nil {
		http.Error(w, "Invalid ID token", http.StatusBadRequest)
		return
	}

	session, err := lineHandler.store.Get(r, "line-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["sub"] = userData["sub"].(string)
	session.Values["name"] = userData["name"].(string)
	session.Values["picture"] = userData["picture"].(string)

	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
