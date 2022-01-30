package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	myNumber, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(myNumber), 64)

	// error handling
	if err != nil {
		panic(err)
	}

	fmt.Println("The number you entered is: ", aFloat)
}
