package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// func main() {
// 	response, err := http.Get("https://example.com")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(body))
// }

func main() {
	go responseSize("https://example.com")
	go responseSize("https://golang.org")
	go responseSize("http://golang.org/doc")
	time.Sleep(3 * time.Second)
}

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}
