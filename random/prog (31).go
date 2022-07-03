package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}

// Map is unordered collection. To get an ordered list of your values. Sort the key first passing
// it to a slice. Sort the slice. Use the sorted slice to get the name and access the value using the map
// with it's key. grades[name]. Grades is the map. Inside it is name which is sorted in the slice

