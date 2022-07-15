// typeEmbedding.go
package main

import . "fmt"

type HelloWorld struct{}

func (h HelloWorld) String() string {
	return "Hello Pinas!"
}

type Message struct {
	HelloWorld
}

func main() {
	m := &Message{}
	Println(m.HelloWorld.String())
	Println(m.String())
	Println(m)
}
