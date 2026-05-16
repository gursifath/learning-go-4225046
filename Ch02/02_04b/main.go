package main

import "fmt"

func main() {

	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum is ", intSum)

	f1, f2, f3 := 4.55, 7.99, 2.87
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum is ", floatSum)

	// total := i1 + f3 // invalid arithmetic operation between mismatching types

	total := float64(i1) + f3
	fmt.Println(total)
}
