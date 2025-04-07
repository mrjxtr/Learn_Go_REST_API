package main

import "fmt"

func main() {
	result := factorial(5)
	fmt.Println(result)
}

func factorial(number int) int {
	// recursion approach
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)

	// for loop approach
	// result := 1
	//
	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }
	//
	// return result
}
