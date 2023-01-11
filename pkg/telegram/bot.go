package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lampadovnikita/LinkScreen/pkg/config"
)

type Bot struct {
	bot           *tgbotapi.BotAPI
	commands      config.Commands
	localizations map[string]config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, commands config.Commands, localizations map[string]config.Messages) *Bot {
	return &Bot{bot: bot, commands: commands, localizations: localizations}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message, err)
			}

			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message, err)
		}
	}
}
