package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	//fmt.Println("Day", dow)

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Sunday"
	case 2:
		result = "Monday"
	case 3:
		result = "Tuesday"
	case 4:
		result = "Wednesday"
	case 5:
		result = "Thursday"
	case 6:
		result = "Friday"
	case 7:
		result = "Saturday"

	}
	fmt.Println(result)
}
