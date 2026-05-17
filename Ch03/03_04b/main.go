package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an array
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// This is a slice (it doesn't have a number or size within square brackets)
	var colorsSlice = []string{"Red", "Green", "Blue"}
	fmt.Println(colorsSlice)

	// more common and memory efficient way to declare slice
	var colorsSliceBetter = make([]string, 0, 4)
	// 0 is the number of initial items in the slice
	// 3 is the capacity of the slice

	fmt.Println(colorsSliceBetter)

	colorsSliceBetter = append(colorsSliceBetter, "Pink", "Orange", "Black", "Maroon")
	fmt.Println(colorsSliceBetter)

	colorsSliceBetter = remove(colorsSliceBetter, 1)
	fmt.Println(colorsSliceBetter)

	sort.Strings(colorsSliceBetter)
	fmt.Println(colorsSliceBetter)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
