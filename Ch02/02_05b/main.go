package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	result := math.Round(sum*100) / 100
	fmt.Printf("Float sum: %.5f\n", result)

	fmt.Println("Value of Pi is ", math.Pi)

	circleRadius := 13.5
	circumference := 2 * circleRadius * math.Pi

	fmt.Printf("The circumference is %.4f\n", circumference)
}
