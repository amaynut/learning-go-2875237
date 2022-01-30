package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("******** A simple calculator *******")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter value1: ")
	input1, _ := reader.ReadString('\n')
	num1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err1 != nil {
		panic("Value1 must be a number")
	}

	fmt.Print("Please enter value2: ")
	input2, _ := reader.ReadString('\n')
	num2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err2 != nil {
		panic("Value2 must be a number")
	}

	sum := num1 + num2

	fmt.Printf("The sum of %v and %v is: %.2f", num1, num2, sum)

}
