package main

import (
	"fmt"
	"os"
)

// using custom types to enhance developer experience
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func makeMaps() floatMap {
	courseRatings := make(floatMap, 2) // only one addition parameter
	fmt.Println(courseRatings)

	courseRatings["TypeScript"] = 5.9
	courseRatings["Vue"] = 5.9

	// fmt.Println(courseRatings)
	courseRatings.output() // using custom methods on type aliases

	return courseRatings
}

func main() {
	exit := true
	courseRatings := makeMaps()

	for k, v := range courseRatings {
		// do something per iteration
		fmt.Printf("key: %v, val: %v\n", k, v)
	}

	if exit {
		fmt.Printf("\nexiting early\n")
		os.Exit(0)
	}
	// Maps are key value pairs like pythong's dictionaries
	emptymap := map[string]string{} // this is an empyty map
	websites := map[string]string{  // this is a map with values
		"Google":              "https://google.com",
		"GitHub":              "https://github.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(emptymap) // prints the empty map
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"]) // accessign a value

	websites["Facebook"] = "https://facebook.com" // adding a new value in the map
	fmt.Println(websites)

	websites["Facebook"] = "facebook.com" // edditing a value
	fmt.Println(websites)

	delete(websites, "Facebook") // deleting a value
	fmt.Println(websites)
}
