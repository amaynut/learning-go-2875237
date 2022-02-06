package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Calculator program
func main() {
	// get input values
	value1 := getInputValue(1)
	value2 := getInputValue(1)

	// get operator
	operator := getInputOperator()

	// make calculation
	result := calculate(operator, value1, value2)

	// format the result
	fmt.Printf("The result is %.2f", result)
}

func calculate(operator string, value1, value2 float64) float64 {
	switch operator {
	case "-":
		return value1 - value2
	case "+":
		return value1 + value2
	case "*":
		return value1 * value2
	case "/":
		return value1 * value2
	default:
		panic("Invalid Operator")
	}
}
func getInputValue(num int) float64 {
	// prompt the user to enter a value
	print("Enter value", num, ": ")

	// read the value
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// convert the input to a number
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic("The value is not a number")
	}

	return value
}

func getInputOperator() string {
	// prompt the user to enter a value
	print("Enter an operator (+, -, *): ")

	// read the value
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// convert the input to a number
	operator := strings.TrimSpace(input)

	allowedOperators := []string{"+", "-", "*", "/"}

	if InArray(operator, allowedOperators) == -1 {
		panic("Invalid Operator, it must be one of the following values: [*, - , +]")
	}

	return operator
}

// InArray check of a given value is in an array
func InArray(val interface{}, array interface{}) (index int) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}
