package main

import (
	"fmt"
)

func main() {

	// declare a nil slice of strings.
	var data []string

	// capture the capacity of the slice.
	lastCap := cap(data)

	// apend ~100k strings to the slice.
	for record := 1; record <= 1e5; record++ {

		// use the built-in function append to add to the slice.
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		// when the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {

			// calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// save the new values for capacity
			lastCap = cap(data)

			// display results
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %.2f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)

		}

	}
}
