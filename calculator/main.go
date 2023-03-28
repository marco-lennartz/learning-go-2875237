package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Value 1:")
	strIn, _ := reader.ReadString('\n')
	v1, err := strconv.ParseFloat(strings.TrimSpace(strIn), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a numeric")
	}

	fmt.Println("Value 2:")
	strIn, _ = reader.ReadString('\n')
	v2, err := strconv.ParseFloat(strings.TrimSpace(strIn), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a numeric")
	}

	result := math.Round((v1+v2)*100) / 100

	fmt.Printf("The sum of %.2f and %.2f is %.2f", v1, v2, result)

}
