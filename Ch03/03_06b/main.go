package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")

	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

	poodle.Weight = 67
	fmt.Println(poodle.Weight)
	fmt.Println(poodle)
}

// capital starting letter signifies the function/field/variable/type as PUBLIC
type Dog struct {
	Breed  string
	Weight int
}
