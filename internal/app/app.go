package app

import (
	"github.com/sandmanitsu/market-bot-telegram/internal/config"
	vkapi "github.com/sandmanitsu/market-bot-telegram/internal/vk-api"
)

func NewApp(cfg *config.Config) {
	vkapi.MarketGet(cfg.VK)
}
