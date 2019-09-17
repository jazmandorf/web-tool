package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

func init() {

}

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob(`./src/views/*.html`)),
	}
	e.Renderer = renderer
	e.Handle
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	e.POST("/hello", func(c echo.Context) error {
		return c.Render(200, "hello.html", map[string]interface{}{
			"hello": 1234,
		})
	})
	e.Logger.Fatal(e.Start(":1234"))

}
