package main

import (
	"fmt"
)

func main() {
	// string array unassigned value will have empty string value
	var colors [8]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])

	// int array unassigned element will have 0 value
	var numbers = [5]int{1, 4, 6, 3}
	fmt.Println(numbers)
	fmt.Println("The number of colors is:", len(colors))
	fmt.Println("The number of numbers is:", len(numbers))
}
