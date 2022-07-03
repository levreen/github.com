// echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(i, s)
		fmt.Println(strings.Join(os.Args[:], sep)) // func strings.Join "concatenate" - exact word
	}

}
