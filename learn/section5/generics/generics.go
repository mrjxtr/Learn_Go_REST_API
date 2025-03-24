package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := add(1, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = add("3", "4")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	result, err = add(5.25, 6.5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

type Number interface {
	int | float64
}

// Make a generic function that takes in integers, floats, or strings
func add[T Number | string](a, b T) (float64, error) {
	var x, y float64
	var err error

	// Type switch:
	// if a is a string, convert it to a float
	// if a is an int, convert it to a float
	switch v := any(a).(type) {
	case string:
		x, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
	case int:
		x = float64(v)
	case float64:
		x = v
	default:
		return 0, fmt.Errorf("unsupported type")
	}

	// Type switch:
	// if b is a string, convert it to a float
	// if b is an int, convert it to a float
	switch v := any(b).(type) {
	case string:
		y, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, err
		}
	case int:
		y = float64(v)
	case float64:
		y = v
	default:
		return 0, fmt.Errorf("unsupported type")
	}

	return x + y, nil
}
