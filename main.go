package main

import (
	"fmt"
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

	// paresGlob 를 사용하여 모든 경로에 있는 파일을 가져올 경우 사용하면 되겠다.
	// 사용한다음에 해당 파일을 불러오면 되네.
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob(`./src/views/*.html`)),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name":    "Dolly!",
			"reverse": 1234,
		})
	}).Name = "foobar"

	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", map[string]interface{}{
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

	e.GET("/getJson", func(c echo.Context) error {
		url := `"http//localhost:1234/getTest?email=jazmandorf@gmail.com&name=Dennis"`

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("에러1")
			log.Fatal(err)
		}

		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("에러2")
			log.Fatal(err)
		}
		fmt.Println("data : ", data)
		return c.String(http.StatusOK, "gethtml")
	})

	e.GET("/getHtml", func(c echo.Context) error {
		url := "http//localhost:1234/getTest"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("에러1")
			log.Fatal(err)
		}

		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("에러2")
			log.Fatal(err)
		}
		fmt.Println("data : ", data)

		//return c.HTML(200, string(data))
		return c.String(http.StatusOK, "gethtml")
	})

	e.Logger.Fatal(e.Start(":1234"))

}

type myStruct struct {
	Name   string
	Age    int
	Height int
}
