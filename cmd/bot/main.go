package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lampadovnikita/LinkScreen/pkg/config"
	"github.com/lampadovnikita/LinkScreen/pkg/telegram"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		panic(err)
	}
	botApi.Debug = true

	bot := telegram.NewBot(botApi, cfg.Commands, cfg.Localizations)
	bot.Start()
}
