package main

import (
	"fmt"
)

func main() {
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
