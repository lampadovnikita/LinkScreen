package telegram

import (
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case b.commands.Start:
		return b.handleStartCommand(message)
	case b.commands.Help:
		return b.handleHelpCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	text := b.localizations[b.getLangCode(message)].Responses.Start
	err := b.sendResponseText(message.From.ID, text)

	return err
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	text := b.localizations[b.getLangCode(message)].Responses.Help
	err := b.sendResponseText(message.From.ID, text)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	text := b.localizations[b.getLangCode(message)].Responses.UnknownCommand
	err := b.sendResponseText(message.From.ID, text)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		return nonURLError
	} else {
		if err = b.handleURL(message.Chat.ID, message.Text); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleURL(chatID int64, url string) error {
	imageBytes, err := TakeScreen(url)
	if err != nil {
		return err
	}

	if err := b.sendPhoto(chatID, imageBytes); err != nil {
		return err
	}

	return nil
}

func (b *Bot) getLangCode(msg *tgbotapi.Message) string {
	return msg.From.LanguageCode
}
