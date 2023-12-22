package libraryloadtemplate

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/labstack/gommon/log"
)

// LoadTemplates is a custom HTML renderer to support multi templates
func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Load layout templates
	layouts, err := filepath.Glob(filepath.Join(templatesDir, "layouts", "*"))
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}

	// Load all templates, excluding layouts
	includes, err := filepath.Glob(filepath.Join(templatesDir, "**", "*"))
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}

	// Generate the templates map from layout and include directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		// Use filepath.ToSlash to ensure consistency across different OS
		r.AddFromFiles(filepath.ToSlash(filepath.Base(include)), files...)
	}

	return r
}
