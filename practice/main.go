package main

import (
	"fmt"
)

func main() {
	anInt := 2658
	p := &anInt
	fmt.Println("The memory address p is pointing to is: ", p)
	fmt.Println("The value of p is: ", *p)

	// change the original variable
	anInt = 555
	fmt.Println("The memory address p is pointing to is still: ", p)
	fmt.Println("The value of p is now: ", *p)

	// change the pointer variable
	*p = 888
	fmt.Println("The value of p is now: ", *p)
	fmt.Println("The value of anInt is now: ", anInt)

}
