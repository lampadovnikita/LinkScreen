package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	nonURLError = errors.New("Text is not a link")
)

func (b *Bot) handleError(message *tgbotapi.Message, err error) {
	var messageText string

	langCode := message.From.LanguageCode

	switch err {
	case nonURLError:
		messageText = b.localizations[langCode].Errors.NonURL
	default:
		messageText = b.localizations[langCode].Errors.Default
	}

	msg := tgbotapi.NewMessage(message.From.ID, messageText)
	b.bot.Send(msg)
}
