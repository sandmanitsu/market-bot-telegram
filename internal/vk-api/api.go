package vkapi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sandmanitsu/market-bot-telegram/internal/config"
)

const (
	api_version = "5.131" // vk api version
)

func MarketGet(cfg config.VK) {
	url := fmt.Sprintf("https://api.vk.com/method/market.get?owner_id=%s&v=%s&access_token=%s", 154529104, api_version, cfg.AccessToken)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("market get error: %s", err)
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
