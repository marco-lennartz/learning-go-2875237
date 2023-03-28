package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered: ", input)

}
