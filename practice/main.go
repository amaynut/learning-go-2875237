package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Functions")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter value 1: ")
	input1, _ := reader.ReadString('\n')
	val1, err1 := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)

	if err1 != nil {
		panic(err1)
	}

	fmt.Print("Enter value 2: ")
	input2, _ := reader.ReadString('\n')
	val2, err2 := strconv.ParseInt(strings.TrimSpace(input2), 10, 64)

	if err2 != nil {
		panic(err2)
	}

	sum := addValues(val1, val2)

	fmt.Println("Sum is: ", sum)

	total, argumentsCount := addAllValues(154, 744, 954, 694, 74)

	fmt.Printf("The total is %v and the number of argument is: %v ", total, argumentsCount)

}

func addValues(val1, val2 int64) int64 {
	return val1 + val2
}

// variable number of arguments and return more than one value
func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
