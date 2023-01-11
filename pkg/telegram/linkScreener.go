package telegram

import (
	"context"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func TakeScreen(url string) (*[]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var imageBuf []byte
	if err := chromedp.Run(ctx, screenshotTask(url, &imageBuf)); err != nil {
		return nil, err
	}

	return &imageBuf, nil
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
