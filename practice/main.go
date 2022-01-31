package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	// in go, a slice is an array with no fixed size
	var sliceColors = []string{"Red", "Green", "Blue"}
	var arrayColors = [3]string{"Red", "Green", "Blue"}
	fmt.Println("The type of sliceColors is: ", reflect.TypeOf(sliceColors))
	fmt.Println("The type of arrayColors is: ", reflect.TypeOf(arrayColors))

	// add an element to an array
	sliceColors = append(sliceColors, "Purple")
	fmt.Println(sliceColors)

	// delete an element from an array (specify the items to keep)
	// delete the first element (shift in php)
	sliceColors = append(sliceColors[1:])
	fmt.Println(sliceColors)

	// delete the last element (pop in php)
	sliceColors = append(sliceColors[:len(sliceColors)-1])
	fmt.Println(sliceColors)

	// declare a slice with an initial size and a capacity (cap)
	numbers := make([]int, 5, 9)
	numbers[1] = 587
	numbers[0] = 5
	numbers[4] = 6
	numbers[3] = 87
	fmt.Println(numbers)
	// sort the slice
	sort.Ints(numbers)
	fmt.Println(numbers)
}
