package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lampadovnikita/LinkScreen/pkg/config"
	"github.com/lampadovnikita/LinkScreen/pkg/telegram"
)

const (
	// The env variable, which stores the full path to the folder with the configuration files, from the project root folder which is ./pkg/config
	configFilesFolderPathEnvName = "LINK_SCREEN_CONFIG_FILES_FOLDER_PATH"
)

func main() {
	configFilesFolderPath := os.Getenv(configFilesFolderPathEnvName)

	cfg, err := config.Init(configFilesFolderPath)
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
