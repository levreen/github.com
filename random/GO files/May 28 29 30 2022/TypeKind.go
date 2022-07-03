package main

import (
	"fmt"
)

const (
	// max integer value on 64 bit architecture.
	maxInt = 9223372036854775807

	// much larger value than int64.
	bigger = 9223372036854775808543522345

	// will not compile
	// biggerInt int64 = 9223372036854775808543522345
)

func main() {
	const (
		A1 = iota // 0 : Start at 0
		B1 = iota // 1 : Increment by 1
		C1 = iota // 2 : Increment by 1
	)

	fmt.Println("1: ", A1, B1, C1)

	const (
		A2 = iota // 0 : Start at 0
		B2        // 1 : Increment by 1
		C2        // 2 : Increment by 1
	)

	fmt.Println("2: ", A2, B2, C2)

	const (
		A3 = iota + 1 // 1 : Start at 0 + 1
		B3            // 2 : Increment by 1
		C3            // 3 : Increment by 1
	)

	fmt.Println("3: ", A3, B3, C3)

	const (
		Ldate         = 1 << iota // 1 : Shift 1 to the left 0. 0000 0001
		Ltime                     // 2 : Shift 1 to the left 1. 0000 0010
		Lmicroseconds             // 4 : Shift 1 to the left 2. 0000 0100
		Llongfile                 // 8 : Shift 1 to the left 3. 0000 1000
		Lshoftfile                // 16: Shift 1 to the left 4. 0001 0000
		LUTC                      // 32: Shift 1 to the left 5. 0010 0000
	)

	fmt.Println("Log: ", Ldate, Ltime, Lmicroseconds, Llongfile, Lshoftfile, LUTC)
}

// untyped constants
const ui = 12345    // kind: integer
const uf = 3.141592 // kind: floating-point

// typed constants still use the constant type system but
// is restricted.
const ti int = 12345        // type: int
const tf float64 = 3.141592 // type: float64

// ./constants.go:XX: constant 1000 overflows uint8
// const myUint8 uint8 = 1000

// constant arithmetic supports different kinds.
// kind promotion is used to determine kind in these scenarios

// variable answer will of type float64.
var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)

// constant third will be of kind floating point
const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

// constant zero will be a kind integer.
const zero = 1 / 3 // KindInt(1) / KindInt(3)

// This is an example of contact arithmetic between typed and
// untyped constants. Must have like types to perform math.
const one int8 = 1
