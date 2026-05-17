package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
	fmt.Println(addValues(1, 3))

	fmt.Println(addAllValues(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// void return val and no params
func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

// As types of the input params is int for both values, don't need to write 'int' twice
func addValues2(value1, value2 int) int {
	return value1 + value2
}

// Multiple Variables
func addAllValues(values ...int) (int, int, float64) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	count := len(values)
	avergae := float64(sum) / float64(count)
	return sum, count, avergae
}
