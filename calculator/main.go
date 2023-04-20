package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	value1 := readNumber("Value 1:")
	value2 := readNumber("Value 2:")
	operator := readOperator(reader)

	result := 0.0
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		result = value1 / value2
	}
	result = math.Round(result*100) / 100
	fmt.Printf("The result of %v %v %v is %v\n\n", value1, operator, value2, result)
}

func readNumber(request string) float64 {
	fmt.Print(request)
	input1, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("%v must be a number", request))
	}
	return value
}

func readOperator(reader *bufio.Reader) string {
	fmt.Print("Operator:")
	input1, _ := reader.ReadString('\n')
	value := strings.TrimSpace(input1)
	if len(value) > 1 {
		panic("Der Operator darf nicht länger als 1 Zeichen sein!")
	}
	return value
}
