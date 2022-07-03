package main

import (
	"fmt"
	"sort"
)

type user struct {
	name    string
	surname string
}

func main() {

	// declare and initialize the map with values.
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// pull the keys from the map.
	var keys []string // creating a new slice
	for key := range users {
		keys = append(keys, key)
	}

	// sort the keys alphabetically.
	sort.Strings(keys)

	// walk through the keys and pull each value from the map.
	for _, key := range keys {
		fmt.Println(key, users[key])
	}

}
