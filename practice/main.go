package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "Hello from go!"
	fmt.Println(aString)
	fmt.Printf("The variable aString type is %T\n", aString)

	var myInt int = 42
	fmt.Println(myInt)

	var defaultInt int
	fmt.Println(defaultInt)

	var myString = "test"
	fmt.Println(myString)
	fmt.Printf("The variable myString type is %T\n", myString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable anotherInt type is %T\n", anotherInt)

	myNextString := "String too"
	fmt.Println(myNextString)
	fmt.Printf("The variable myNextString type is %T\n", myNextString)

	fmt.Println(aConst)
	fmt.Printf("The variable aConst type is %T\n", aConst)
}
