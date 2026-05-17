package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	// in Go, we don't need break statements after each case; it's happen automatically
	// however, can use fallthrough if you want to execute the following cases
	var result string
	switch dayNumber = 5; dayNumber {
	case 0:
		result = "It's a Sunday"
	case 1:
		result = "It's a Monday"
	case 2:
		result = "It's a Tuesday"
	case 3:
		result = "It's a Wednesday"
	default:
		result = "None"
	}
	fmt.Println(result)

	x := -65
	switch {
	case x < 0:
		result = "Lesser than 0"
	case x == 0:
		result = "Equal to 0"
		fallthrough // read above for explanation on fallthrough
	default:
		result = "Greater than 0"
	}

	fmt.Println(result)
}
