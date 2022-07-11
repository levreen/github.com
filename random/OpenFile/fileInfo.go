package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Mode())
}
