package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")

	// read about make vs new functions
	// new declares the variables but doesn't allocate memory
	// make() allocates memory to the newly created variable
	states := make(map[string]string)
	states["UP"] = "Lucknow"
	states["Tamil Nadu"] = "Chennai"
	states["Maharashtra"] = "Mumbai"
	fmt.Println(states)

	fmt.Println(states["UP"])
	fmt.Println(states["Maharashtra"])
	up := states["UP"]
	fmt.Println(up)

	delete(states, "UP")
	fmt.Println(states)

	states["Goa"] = "Panaji"
	fmt.Println(states)

	// loop through map
	for k, v := range states {
		fmt.Printf("%s: %s\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("Sorted")

	for i := range keys {
		fmt.Println(keys[i], states[keys[i]])
	}
}
