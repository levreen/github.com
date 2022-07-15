// webApp48
package main

import (
	. "fmt"
	"net"
)

func main() {
	if listener, e := net.Listen("tcp", ":1024"); e == nil {
		for {
			if connection, e := listener.Accept(); e == nil {
				go func(c net.Conn) {
					defer c.Close()
					Fprintln(c, "Hello Pinas!")
				}(connection)
			}
		}
	}
}

/*
$ go run webApp48.go &
$ telnet localhost 1024
*/
