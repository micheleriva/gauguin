package chromium

import (
	"context"
	"fmt"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// GenerateImage returns a byte array representing the generated image
func GenerateImage(html string, width float64, height float64) []byte {
	var buf []byte

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if err := chromedp.Run(ctx, takeScreenshot(&buf, html, width, height)); err != nil {
		panic(err)
	}

	return buf
}

func takeScreenshot(res *[]byte, html string, width float64, height float64) chromedp.Tasks {
	var err error

	return chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("data:text/html,%s", html)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			*res, err = page.CaptureScreenshot().
				WithQuality(100).
				WithClip(&page.Viewport{
					Width:  width,
					Height: height,
					Scale:  1,
				}).Do(ctx)

			if err != nil {
				return err
			}

			return nil
		}),
	}
}
