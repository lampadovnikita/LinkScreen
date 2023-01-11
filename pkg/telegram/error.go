package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	nonURLError = errors.New("Text is not a link")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case nonURLError:
		messageText = b.messages.Errors.NonURL
	default:
		messageText = b.messages.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
