package telegram

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	nonURLError = errors.New("Text is not a link")
)

func (b *Bot) handleError(message *tgbotapi.Message, err error) error {
	log.Println("Handle error:", err)

	var messageText string

	langCode := message.From.LanguageCode

	switch err {
	case nonURLError:
		messageText = b.localizations[langCode].Errors.NonURL
	default:
		messageText = b.localizations[langCode].Errors.Default
	}

	if err := b.sendResponseText(message.From.ID, messageText); err != nil {
		return err
	}

	return nil
}
