package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")

	// someInt := 42 (can also directly define variable in if condition)
	var result string

	if someInt := 42; someInt < 0 {
		result = "lesser than"
	} else if someInt > 0 {
		result = "greater than"
	} else {
		result = "equal to"
	}
	fmt.Println(result)
}
