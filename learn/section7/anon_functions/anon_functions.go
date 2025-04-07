package main

import "fmt"

// anonymous functions are useful if you have a function that wants a function
// or if you have a function that returns a function
func main() {
	numbers := []int{1, 2, 3}

	// passing in a function definition and logic here
	// instead of creating the fucntions ahead of time
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)

	// Example of using anonymus fucntions with "closures"
	// Creating function variables using the "Factory Function" createTransformer()
	double := createTransformer(2)
	triple := createTransformer(3)

	// using the created fucntions as a parameter here
	doubledN := transformNumbers(&numbers, double)
	tripledN := transformNumbers(&numbers, triple)

	fmt.Println(doubledN)
	fmt.Println(tripledN)
}

// this is the fucntions that expects a fucntion as a parameter
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// This is a what's called a "Factory Function"
// This is also a "closure" which locks the "factor" value into
// the returned result of createTransformer() which makes
// available to call again later in codebase
func createTransformer(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}
