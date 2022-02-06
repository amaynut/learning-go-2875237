package main

import (
	"fmt"
	"strings"
)

func main() {
	poodle := Dog{"Poodle", 10, "Haw haw!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.SpeakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak Type method (exported)
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes is how to dog speaks loudly
func (d Dog) SpeakThreeTimes() {
	fmt.Println(strings.Repeat(d.Sound+" ", 3))
}
