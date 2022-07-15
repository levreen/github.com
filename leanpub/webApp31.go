// webApp31.go
package main

import (
	. "fmt"
	. "net/http"
)

const MESSAGE = "hello world"
const ADDRESS = ":1025"

func main() {
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Context-Type", "text/plain")
		Fprintf(w, MESSAGE)
	})
	ListenAndServe(ADDRESS, nil)
}
