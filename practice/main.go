package main

import (
	"fmt"
)

func main() {
	anInt := 11
	var p = &anInt
	fmt.Println("Value if p:", *p)

	value1 := 42.42
	pointer1 := &value1
	fmt.Println("Value if pointer1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value if pointer1:", *pointer1)
	fmt.Println("Value if value1:", value1)

}
