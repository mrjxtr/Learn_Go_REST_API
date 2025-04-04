package main

import (
	"fmt"
	"os"
)

func dynamicSlice() {
	emptySlice := []float64{} // creating an empty slice
	exampleSlice := []float64{1.2, 3.4}
	fmt.Println(emptySlice)
	fmt.Println(exampleSlice)

	// assigning values to slice
	exampleSlice[1] = 5.6
	// exampleSlice[2] = 7.8 // this is not allowed since the slice only has 2 values
	// instead we can do this:
	updatedleSlice := append(exampleSlice, 7.8)
	fmt.Println(exampleSlice, updatedleSlice)
}

func main() {
	exitEarly := true

	dynamicSlice()

	// comment this to run all code bellow
	if exitEarly {
		fmt.Printf("\nexiting early\n")
		os.Exit(0)
	}

	// arrays are collection of data with the same type i.e different numbers
	// var emptyArray [4]string    // creating an empty string
	emptyArray := [4]string{}   // better way of creating an empty string
	emptyArray[2] = "Something" // adding a value to a specific index of an array

	fmt.Println(emptyArray) // We have empty values plus "Something" in idx 2

	producNames := [4]string{"Coffee", "Iced Tea", "Milk Tea", "Water"}

	fmt.Println(producNames)
	fmt.Println(producNames[0]) // Print specific elements in array

	prices := [4]float64{10.99, 11.5, 9.5, 5.1}
	priceSlice := prices[1:]  // creating a slice of an array
	priceSlice2 := prices[:3] // creating a slice of an array

	fmt.Println("price:", prices)
	fmt.Println("price slice:", priceSlice)
	fmt.Println("price slice2:", priceSlice2)
	fmt.Println(
		"price slice2 with idx:",
		priceSlice2[2:], // Note that you cannot use negative indexing here
	) // Printing only a portion of a slice

	// modifying an element in a slice also modifies the original array
	// we can see here how slices are just windows to the array
	newPrices := [4]float64{12.99, 13.5, 14.5, 15.15}
	fmt.Println("newPrices", newPrices)

	featuredPrices := newPrices[1:]
	featuredPrices[0] = 1111
	fmt.Println("featuredPrices", featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println("highlightedPrices", highlightedPrices)

	// Slices Metadata
	fmt.Println("len of featuredPrices:",
		len(featuredPrices),
	) // number of values in a slice
	fmt.Println(
		"cap of featuredPrices:",
		cap(featuredPrices),
	) // capacity of the slice which is 3

	fmt.Println("len of highlightedPrices:",
		len(highlightedPrices),
	)
	fmt.Printf(
		`cap of highlightedPrices: %v

Cap is 3 since we have more items TO THE RIGHT of
featuredPrices (which is a slice of newPricews) and
highlightedPrices a slice of featuredPrices.
We do not see 4 because the cap does not look at
values on the left of newPricews (the original array) 
`,
		cap(highlightedPrices),
	)

	// in this example, since we are looking at everything
	// to the left of highlightedPrices, cap is also 3
	highlightedPrices = highlightedPrices[:3] // this shows you there are more items to the right
	fmt.Printf("\nnew highlightedPrices %v\n", highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
