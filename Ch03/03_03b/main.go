package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)

	var numbers = [5]int{4, 6, 8, 0, 4}
	fmt.Println(numbers)
	fmt.Println(numbers[4])

	fmt.Printf("Length of colors is %d\n", len(colors))
	fmt.Printf("Length of numbers is %d\n", len(numbers))
}
