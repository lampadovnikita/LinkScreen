package telegram

import (
	"context"
	"net/url"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
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
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var imageBuf []byte
	if err := chromedp.Run(ctx, screenshotTask(url, &imageBuf)); err != nil {
		return err
	}

	photoBytes := tgbotapi.FileBytes{Name: "screen", Bytes: imageBuf}

	photo := tgbotapi.NewPhoto(chatID, photoBytes)

	_, err := b.bot.Send(photo)
	if err != nil {
		return err
	}

	return nil
}

func screenshotTask(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(c context.Context) (err error) {
			*imageBuf, err = page.CaptureScreenshot().WithQuality(90).Do(c)

			return err
		}),
	}
}

func (b *Bot) sendResponseText(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getLangCode(msg *tgbotapi.Message) string {
	return msg.From.LanguageCode
}
