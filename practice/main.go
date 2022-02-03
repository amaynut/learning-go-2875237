package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("************ Conditional logic ****************")
	// get a number from the console
	reader := bufio.NewReader(os.Stdin)
	print("Enter a number: ")
	input, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic("The value is not a number")
	}

	// use an if else statement
	var result string
	if number < 0 {
		result = "It's a negative number"
	} else if number == 0 {
		result = "Equal to zero"
	} else {
		result = "It's a positive number"
	}
	println(result)

	// declare a variable inside the if condition
	if currentHour := time.Now().Hour(); currentHour < 7 {
		fmt.Print("Good Night")
	} else if currentHour >= 7 && currentHour < 12 {
		print("Good morning")
	} else if currentHour >= 12 && currentHour < 18 {
		print("Good afternoon!")
	} else {
		print("Good evening")
	}

}
