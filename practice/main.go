package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"
	fmt.Println("Arrays", colors)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("numbers", numbers)
	fmt.Println("No of numbers", len(numbers))
}
