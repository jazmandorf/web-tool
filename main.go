package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func init() {

}

type user struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
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
	e.Static("/assets", "./src/static/assets")
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob(`./src/views/*.html`)),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name":    "Dolly!",
			"reverse": 1234,
			"foobar":  "",
		})
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello.html", map[string]interface{}{
			"Name": myStruct{Name: "Dennis", Age: 36, Height: 170},
		})
	})

	e.GET("/dashboard", func(c echo.Context) error {
		return c.Render(http.StatusOK, "dashboard_2.html", map[string]interface{}{
			"Name": myStruct{Name: "Dennis", Age: 36, Height: 170},
		})
	})

	e.GET("/initial", func(c echo.Context) error {
		return c.Render(http.StatusOK, "form_wizard.html", map[string]interface{}{})
	})

	e.GET("/dashboard2", func(c echo.Context) error {
		return c.Render(http.StatusOK, "dashboard_3.html", map[string]interface{}{
			"Name": myStruct{Name: "Dennis", Age: 36, Height: 170},
		})
	})

	e.POST("/testPost", func(c echo.Context) error {
		return c.String(http.StatusOK, "testPost")
	})

	e.GET("/getTest", func(c echo.Context) error {
		u := new(user)
		if err := c.Bind(u); err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/getHtml", func(c echo.Context) error {
		res, err := http.NewRequest("GET", "http://localhost:1234/initial", nil)
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		return c.HTML(200, string(data))
	})

	e.Logger.Fatal(e.Start(":1234"))

}

type myStruct struct {
	Name   string
	Age    int
	Height int
}
