package main

import "fmt"

// Variadic Fucntions is a function that can work
// with any amount of parameters
func main() {
	numbers := []int{0, 1, 10, 15, 20, 30}

	sum := sumup(0, 1, 10, 15, 20, 30)
	anotherSum := sumup(0, numbers...) // use "..." to unpack list values

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// the "..." collects all the values passed and merged
// it automatically into a slice.
func sumup(start int, numbers ...int) int {
	sum := start

	for _, v := range numbers {
		sum += v
	}

	return sum
}
