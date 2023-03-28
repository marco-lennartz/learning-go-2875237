package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Dates and times", n)

	t := time.Date(2023, time.March, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Dates and times", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Mar 10 23:00:00 2023")
	fmt.Printf("%T", parsedTime)
}
