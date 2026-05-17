package main

import (
	"fmt"
	"time"
)

func main() {
	// turn any function to a goroutine by prepending go keyword
	go say("hello from Amsterdam!!!")
	fmt.Println("Goroutines")

	// anonymous go routines
	go func(message string) {
		fmt.Println(message)
	}("Hello from the Anonymous function")

	time.Sleep(2 * time.Second)
	fmt.Println("All Done!!")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
