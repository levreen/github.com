package main

import (
	"fmt"
)

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	v := 1
	y := incr(&v)
	incr(&y)              // bago e run val of y is 2 tas passing y sa func incr it increases the value by 1
	fmt.Println(incr(&y)) // y := incr(&v)

}
