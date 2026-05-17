package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")
	// use * to refer to the value of the pointer
	// use & to point to the memory address 

	anyInt := 42
	var p *int = &anyInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value of p is", *p)
	}

	anyFloat := 42.13
	pointer1 := &anyFloat

	fmt.Println("Value of pointer1 is", *pointer1)
	*pointer1 = *pointer1 / 13
	fmt.Println("New value is", *pointer1)
	fmt.Println("Original value is", anyFloat)
}
