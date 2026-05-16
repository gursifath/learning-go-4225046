package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	fmt.Println("Hello from Go!")
	fmt.Println(str1, str2, str3)
	strLength, err := fmt.Println("The number is", aNumber)
	if err == nil {
		fmt.Println("The length of string is", strLength)
	}

	fmt.Printf("The value is %v\n", aNumber)
	fmt.Printf("The type is %T\n", aNumber)
	fmt.Printf("The type is %T\n%T\n%T\n", str1, str2, str3)
}
