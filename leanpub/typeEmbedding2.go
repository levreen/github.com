// typeEmbedding2.go
package main

import . "fmt"

type HelloWorld bool

func (h HelloWorld) String() (r string) {
	if h {
		r = "Hello Pinas!!!"
	}
	return
}

type Message struct {
	HelloWorld
}

func main() {
	m := &Message{HelloWorld: true}
	Println(m)
	m.HelloWorld = false
	Println(m)
	m.HelloWorld = true
	Println(m)
}
