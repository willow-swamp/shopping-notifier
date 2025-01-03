package models

type LoginUser struct {
	Sub     string `json:"sub"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
