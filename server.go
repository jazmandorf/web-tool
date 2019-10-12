package main

import (
	"log"
	"os"
	"text/template"
)

type Student struct {
	Name   string
	Age    int
	Emails []string
}

const tmpl = `{{$name := .Name}}
The name is {{$name}}.
{{range .Emails}}
	Myname is {{$name}} email id is {{.}}
{{end}}
`

func main() {
	// url := "http://localhost:1234/getTest"

	// proxyReq, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := &http.Client{}
	// proxyRes, err := client.Do(proxyReq)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer proxyRes.Body.Close()

	// bytes, _ := ioutil.ReadAll(proxyRes.Body)
	// str := string(bytes)
	// fmt.Println(str)

	s := Student{"Dennis", 32, []string{"jazmandorf", "thaeao", "wlqkrdl"}}

	t := template.New("person template")

	t, err := t.Parse(tmpl)

	if err != nil {
		log.Fatal(err)
	}

	err1 := t.Execute(os.Stdout, s)

	if err1 != nil {
		log.Fatal(err1)
	}

}
