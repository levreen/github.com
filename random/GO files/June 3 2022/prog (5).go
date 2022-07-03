package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// declare a string with both chinese and english characters.
	s := "世界 means world"

	// UTFMax is 4 -- up to 4 bytes per encoded rune.
	var buf [utf8.UTFMax]byte

	// iterate over the string.
	for i, r := range s {

		// capture the number of bytes for this rune.
		rl := utf8.RuneLen(r)

		// calculate the slice offset for the bytes associated
		// with this rune.
		si := i + rl

		// copy of rune from string to our buffer

		// destination: the buffer slice itself
		// source: create new string value that starts at index position
		// i and goes for rl (e bytes for the Chinese characters,
		// 1 bye for each English char)

		copy(buf[:], s[i:si])

		// display details
		fmt.Printf("%2d: %q; codepoint:\t%#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
