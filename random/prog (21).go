package main

import (
	"fmt"
	"net"
)
// go type identity - pag pareho sila ng element? sa slice at array yon 
// yung statement nila ay pareho. If then, pweding magpalit na values tulad ng code sa baba
func main() {
	var foo net.IP
	var bar []byte = []byte{1, 2, 3, 4}
	foo = bar

	fmt.Println(foo)
}
