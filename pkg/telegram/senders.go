package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) sendPhoto(chatID int64, imageBytes *[]byte) error {
	photoBytes := tgbotapi.FileBytes{Name: "screen", Bytes: *imageBytes}

	photo := tgbotapi.NewPhoto(chatID, photoBytes)

	_, err := b.bot.Send(photo)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) sendResponseText(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.bot.Send(msg)

	return err
}
