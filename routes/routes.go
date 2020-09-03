package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateViews struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateViews) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Index() *echo.Echo {
	e := echo.New()
	renderer := &TemplateViews{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	// Named route "foobar"
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "INDEX",
		})
	})

	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", map[string]interface{}{
			"name": "ABOUT",
		})
	})

	return e
}
