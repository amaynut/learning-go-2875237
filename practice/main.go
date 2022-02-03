package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// regular for loop
	for i := 0; i < len(colors); i++ {
		println(colors[i])
	}

	// use range (similar to foreEach)
	println("\n************** use range in for loop *********\n")
	for _, color := range colors {
		println(color)
	}

	// equivalent while loop
	value := 1
	for value < 100 {
		println(value)
		if value == 10 {
			break
		}
		value++
	}

	// use a goto statement and a label
	sum := 1
	for sum < 1000 {
		sum += sum
		println("Sum: ", sum)
		if sum >= 200 {
			goto theEnd
		}
	}

theEnd:
	println("End of the Program")
}
