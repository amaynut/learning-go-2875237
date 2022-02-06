package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	//fmt.Print(content)

	// loop through

	tours := toursFromJson(content)
	for i, tour := range tours {
		fmt.Printf("Tour %v: %v. It costs $%v\n", i, tour.Name, tour.Price)
	}
}

// decode the json data
func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))

	// check if it's valid JSON
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour

	for decoder.More() {
		// decode the json value and stored in the pointer
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}

		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
