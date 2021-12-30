package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	chromium "github.com/micheleriva/gauguin/chromium"
	config "github.com/micheleriva/gauguin/config"
)

// ImageSize represents the size of a given image
type ImageSize struct {
	width  float64
	height float64
}

type WebHook struct {
	Action       string `json:"action"`
	Route        string `json:"route"`
	Template     string `json:"template"`
	CacheControl string `json:"cache-control"`
	Size         string `json:"size"`
	Params       map[string]string `json:"params"`
}

// ConfigError shows an error at frontend if configuration has some errors
func ConfigError(c *gin.Context) {
	c.Data(
		http.StatusOK,
		"text/html; charset=utf-8",
		[]byte(`
			<h1>Error</h1>
			<p>
				An error occurred while trying to read <b>Gauguin</b> configuration. <br />
				You can find more useful information in your server logs.
			</p>
		`),
	)
}

// HandleRoutes handles all the Gauguin routes
func HandleRoutes(c *gin.Context) {
	var err error
	params := make(map[string]string)
	route := getCurrentRouteConfig(c)

	for _, param := range route.Params {
		params[param] = c.Query(param)
	}

	templateString, err := ioutil.ReadFile(route.Template)
	if err != nil {
		panic(err)
	}

	t := template.New(route.Template)
	t, err = t.Parse(string(templateString))
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, params); err != nil {
		panic(err)
	}

	isDev := c.Query("dev") != ""

	if isDev {
		c.Data(http.StatusOK, "text/html; charset=utf-8", tpl.Bytes())
		return
	}

	sizes := getImageSize(route.Size)

	image := chromium.GenerateImage(tpl.String(), sizes.width, sizes.height)
	img := bytes.NewReader(image)

	conf := config.Config

	if conf.Hooks.BeforeResponse.Url != "" && !isDev {
		callHook(conf.Hooks.BeforeResponse, "beforeResponse", route, params)
	}

	if route.CacheControl != "" {
		c.Header("Cache-Control", fmt.Sprintf(route.CacheControl))
	}
	c.Render(http.StatusOK, render.Reader{ContentType: "image/jpeg", ContentLength: int64(img.Len()), Reader: img})

	if conf.Hooks.AfterResponse.Url != "" && !isDev {
		callHook(conf.Hooks.AfterResponse, "afterResponse", route, params)
	}
}

func getCurrentRouteConfig(c *gin.Context) config.ConfigV001Route {
	path := c.Request.URL.Path

	for _, route := range config.Config.Routes {
		if route.Path == path {
			return route
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"error": fmt.Sprintf("Cannot find path %s in configuration file", path),
	})

	return config.ConfigV001Route{}
}

func getImageSize(str string) ImageSize {
	sizes := strings.Split(str, "x")

	width, err := strconv.ParseFloat(sizes[0], 64)
	if err != nil {
		panic(err)
	}

	height, err := strconv.ParseFloat(sizes[1], 64)
	if err != nil {
		panic(err)
	}

	return ImageSize{
		width:  width,
		height: height,
	}
}

func callHook(hook config.ResponseHook, action string, confRoute config.ConfigV001Route, params map[string]string) {
	client := &http.Client{}

	body, _ := json.Marshal(WebHook{
		Action: action,
		Route: confRoute.Path,
		Template: confRoute.Template,
		CacheControl: confRoute.CacheControl,
		Size: confRoute.Size,
		Params: params,
	})

	responseBody := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", hook.Url, responseBody)
	if err != nil {
		fmt.Printf("Unable to POST hook to %s. Reason: %s", hook.Url, err)
	}

	req.Header.Set("Content-Type", "application/json")

	for _, header := range hook.Headers {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Unable to POST hook to %s. Reason: %s", hook.Url, err)
		return
	}
	fmt.Println(resp.Status)
}