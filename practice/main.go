package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("The sum of the numbers is: ", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	fSum := f1 + f2 + f3
	fmt.Printf("The sum of the floats is: %.2f\n", fSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("The circle circumference is %.2f\n", circumference)

	aF := 35.5689

	fmt.Printf("The rounded float is %.3f\n", aF)

}
