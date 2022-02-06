package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// read an write files
func main() {

	content := "Writing and reading files with Go!"
	fileName := "./fromString.txt"
	// create a file
	file, err := os.Create(fileName)
	checkError(err)
	// write to it
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	// wait (defer) until the write operation is done then close the file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// read the file
	defer readFile(fileName)
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	println("The text in the file is: ", string(data))
}
func checkError(err error) {
	if err != nil {
		println("An error occurred")
		panic(err)
	}
}
