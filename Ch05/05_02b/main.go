package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
	dog.Sound = "Arf"
	dog.Speak()
	fmt.Println(dog.SpeakThreeTimes())
}

type Dog struct {
	Breed string
	Sound string
}

// in Go, a method is a member of a type
func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}

func (d Dog) SpeakThreeTimes() string {
	return fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
}
