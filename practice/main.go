package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 34, 56
	intSum := i1 + i2 + i3

	fmt.Println("int sum is: ", intSum)

	f1, f2, f3 := 12.1, 34.2, 56.3
	floatSum := f1 + f2 + f3

	fmt.Println("float sum is: ", floatSum)
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("float sum is: ", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("circumference: %.2f\n", circumference)
}
