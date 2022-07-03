package main

func AcceptAnything(thing interface{}) {
}

func main() {
	AcceptAnything(3.1415)
	AcceptAnything("A String")
	AcceptAnything(true)
}
