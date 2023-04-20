package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum := addValues(1, 3)
	fmt.Println(sum)
	multiSum, multiCount := addAllValues(1, 3, 5, 6, 234, 21236)
	fmt.Println("Sum of values", multiSum)
	fmt.Println("Counted values", multiCount)
}

func doSomething() {
	fmt.Println("doing something")
}

func addValues(v1 int, v2 int) int {
	return v1 + v2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
