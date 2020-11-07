package editor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	config "github.com/micheleriva/gauguin/config"

	"github.com/gin-gonic/gin"
)

type TemplateData struct {
	GauguinConfig string
}

type StaticAssets struct {
	CSS []string
	JS  []string
}

func RenderEditor(c *gin.Context) {

	var tmplData TemplateData

	conf, err := json.Marshal(config.Config)
	if err != nil {
		log.Fatal(err)
	}

	tmplData.GauguinConfig = string(conf)
	c.HTML(http.StatusOK, "editor.tmpl", tmplData)
}

func GetEditorStaticFiles() StaticAssets {
	var jsFiles []string
	var cssFiles []string

	root := "editor-webapp"

	if err := filepath.Walk(fmt.Sprintf("%s/js", root), func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".js") {
			jsFiles = append(jsFiles, path)
		}
		return err
	}); err != nil {
		log.Fatal(err)
	}

	if err := filepath.Walk(fmt.Sprintf("%s/css", root), func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".css") {
			cssFiles = append(cssFiles, path)
		}
		return err
	}); err != nil {
		log.Fatal(err)
	}

	return StaticAssets{
		CSS: cssFiles,
		JS:  jsFiles,
	}
}
