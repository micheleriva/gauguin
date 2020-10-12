package controller

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	chromium "github.com/micheleriva/gauguin/chromium"
	conf "github.com/micheleriva/gauguin/config"
)

// HandleRoute handles all the Gauguin routes
func HandleRoute(c *gin.Context, route conf.ConfigV001Routes) {
	var err error
	params := make(map[string]string)

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

	image := chromium.GenerateImage(tpl.String(), 1200, 630)
	img := bytes.NewReader(image)

	c.Render(http.StatusOK, render.Reader{ContentType: "image/jpeg", ContentLength: int64(img.Len()), Reader: img})
}
