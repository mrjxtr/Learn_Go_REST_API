package main

import (
	"fmt"
)

// using custom function type definition
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}

	// add only the function name and do not execute with "()"
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	// getting appropriate transformer based on slice input
	transformerFn1 := getTransformerFn(&numbers)
	transformerFn2 := getTransformerFn(&moreNumbers)

	// since transformerFn1 and transformerFn2 are functions, we can use them here
	TNums := transformNumbers(&numbers, transformerFn1)
	moreTNums := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(TNums)
	fmt.Println(moreTNums)
}

// taking in a function "transform" as a "value" or parameter
// by doing this, we simplify this main fucntion instead of calling
// each helper fucntion "double()", "triple()" inside,
// we just pass it in when calling transformNumbers()
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// Returning a function as a result of a fucntion
func getTransformerFn(n *[]int) transformFn {
	if (*n)[0] == 1 {
		return double
	} else {
		return triple
	}
}
