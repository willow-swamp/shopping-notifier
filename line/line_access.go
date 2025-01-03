package line

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type lineAccess struct {
	accessToken string
}

func NewLineAccess(accessToken string) *lineAccess {
	return &lineAccess{
		accessToken: accessToken,
	}
}

func (lineAccess *lineAccess) GetUserData() map[string]interface{} {
	data := url.Values{}
	data.Add("id_token", lineAccess.accessToken)
	data.Add("client_id", os.Getenv("LINE_CHANNEL_ID"))

	resp, err := http.PostForm("https://api.line.me/oauth2/v2.1/verify", data)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println(err)
		return nil
	}

	return result
}
