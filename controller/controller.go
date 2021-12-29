package controller

import (
	"bytes"
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

	c.Header("Cache-Control", "max-age=604800")
	c.Render(http.StatusOK, render.Reader{ContentType: "image/jpeg", ContentLength: int64(img.Len()), Reader: img})
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
