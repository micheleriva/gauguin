package chromium

import (
	"context"
	"fmt"
	"math"

	"github.com/chromedp/cdproto/emulation"
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

	return chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("data:text/html,%s", html)),
		chromedp.ActionFunc(func(ctx context.Context) error {

			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			w, h := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			err = emulation.SetDeviceMetricsOverride(w, h, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			*res, err = page.CaptureScreenshot().
				WithQuality(100).
				WithClip(&page.Viewport{
					X:      0,
					Y:      0,
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
