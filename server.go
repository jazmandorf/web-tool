package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:1234/getTest"

	proxyReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	proxyRes, err := client.Do(proxyReq)
	if err != nil {
		log.Fatal(err)
	}
	defer proxyRes.Body.Close()

	bytes, _ := ioutil.ReadAll(proxyRes.Body)
	str := string(bytes)
	fmt.Println(str)

}
