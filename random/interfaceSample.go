package main

import "fmt"

type Speak interface {
	SayHello() string
}

type English struct{}

func (e English) SayHello() string {
	return "Hello"
}

type Spanish struct{}

func (s Spanish) SayHello() string {
	return "Hola"
}

func NewVoice(lang string) Speak {
	switch lang {
	case "Spanish":
		return Spanish{}

	default:
		return English{}
	}
}

func main() {
	var voice Speak
	voice = NewVoice("Spanish")
	fmt.Println(voice.SayHello())
	voice = NewVoice("Filipino")
	fmt.Println(voice.SayHello())
}
