package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")
	t := time.Date(2026, time.November, 10, 21, 0, 0, 0, time.UTC)
	fmt.Printf("The time Go was released is %s\n", t)

	now := time.Now()
	fmt.Printf("The current time is %s\n", now)
	fmt.Printf("The current time type is %T\n", now)
	fmt.Println("The ANSIC format is", now.Format(time.ANSIC))

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomrorow's time is %s\n", tomorrow)

	// The date to remember is: Jan 2 15:04:05 2006 MST
	// mnemonic in Go is         1  2  3  4  5   6   -7
	// 1 Month
	// 2 Date
	// 3 Hour
	// 4 Minute
	// 5 Second
	// 6 Year
	// -7 Timezone
	newFormat := "Mon 2006-01-02"
	fmt.Printf("Today's Date in newFormat is %s\n", tomorrow.Format(newFormat))

	anotherNewformat := "Mon Jan 2 15:04:05 2006 MST"
	fmt.Printf("Another day is %s\n", now.Format(anotherNewformat))
}
