package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate a random number between 1 and 7
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	// in golang, the break is the default behavior (to execute the next case we need fallthrough keyword
	switch dow {
	case 1:
		print("It's Sunday")
		// fallthrough
	case 2:
		print("It's Monday")
		// fallthrough
	default:
		print("It's some other day")
	}
}
