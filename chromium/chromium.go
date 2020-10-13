package chromium

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/go-resty/resty/v2"
)

// ChromeDevToolsVersion describes the response of google devtool protocol
type ChromeDevToolsVersion struct {
	Browser              string `json:"browser"`
	ProtocolVersion      string `json:"protocol-version"`
	WebKitVersion        string `json:"WebKit-Version"`
	V8Version            string `json:"V8-Version"`
	UserAgent            string `json:"User-Agent"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

var isDockerized bool

func init() {
	isDockerized = os.Getenv("DOCKERIZED") == "true"
}

func getCDTData() ChromeDevToolsVersion {
	var CDTData ChromeDevToolsVersion

	client := resty.New()

	resp, err := client.R().SetHeader("HOST", "localhost").Get("http://alpine_chrome:9222/json/version")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Body(), &CDTData)
	if err != nil {
		log.Fatal(err)
	}

	CDTData.WebSocketDebuggerURL = strings.Replace(CDTData.WebSocketDebuggerURL, "localhost", "alpine_chrome:9222", 1)

	return CDTData
}

// GenerateImage returns a byte array representing the generated image
func GenerateImage(html string, width float64, height float64) []byte {
	var buf []byte
	chromeDevToolsWS := getCDTData()

	allocatorCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), chromeDevToolsWS.WebSocketDebuggerURL)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()

	if err := chromedp.Run(ctx, takeScreenshot(&buf, html, width, height)); err != nil {
		log.Fatal(err)
	}

	return buf
}

func takeScreenshot(res *[]byte, html string, width float64, height float64) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.Navigate(fmt.Sprintf("data:text/html,%s", `<div id="gauguin-root"></div>`)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, exp, err := runtime.Evaluate(fmt.Sprintf("document.write(`%s`)", html)).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil

		}),
		chromedp.ActionFunc(func(ctx context.Context) error {

			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}

			w, h := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			err = emulation.SetDeviceMetricsOverride(w, h, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypeLandscapeSecondary,
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
