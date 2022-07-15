// webApp43
package main

import (
	. "fmt"
	. "net/http"
	"os"
	"sync"
)

const SECURE_ADDRESS = ":1025"

var address string
var servers sync.WaitGroup

// You can set the HTTP port
// $ SERVE_HTTP=":3030" go run 43.go
func init() {
	// Declaration, Condition
	if address = os.Getenv("SERVE_HTTP"); address == "" { // If statement initialization.
		address = ":1024"
	}
}

func main() {
	message := "hello pinas!"
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	Launch(func() {
		ListenAndServe(address, nil)
	})

	Launch(func() {
		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	})
	servers.Wait()
}

func Launch(f func()) {
	servers.Add(1)
	go func() {
		defer servers.Done()
		f()
	}()
}
