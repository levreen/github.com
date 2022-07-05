package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

// func main() {
// 	sizes := make(chan int)
// 	go responseSize("https://example.com", sizes)
// 	go responseSize("https://golang.org", sizes)
// 	go responseSize("http://golang.org/doc", sizes)
// 	fmt.Println(<-sizes)
// 	fmt.Println(<-sizes)
// 	fmt.Println(<-sizes)
// }
type Page struct {
	URL  string
	Size int
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
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
	channel <- Page{URL: url, Size: len(body)}
}
