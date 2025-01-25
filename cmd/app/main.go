package main

import (
	"log"
	"log/slog"

	"github.com/sandmanitsu/market-bot-telegram/internal/app"
	"github.com/sandmanitsu/market-bot-telegram/internal/config"
	"github.com/sandmanitsu/market-bot-telegram/internal/logger"
)

func main() {
	log.Println("config initializing...")
	config := config.MustLoad()

	log.Println("logger initializing...")
	logger := logger.NewLogger(config.Application.Env)
	logger.Info("logger started!", slog.String("env", config.Application.Env))

	// ??? here add storage

	logger.Info("Starting app...")
	app.NewApp(config)
}
