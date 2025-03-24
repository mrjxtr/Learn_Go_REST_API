package main

import (
	"fmt"
)

type something struct {
	text string
}

func main() {
	s := something{}

	printSomething(1)       // can print int
	printSomething(1.5)     // can print float
	printSomething("Hello") // can print string
	printSomething(s)       // can take in struct but will not do anything
}

// i: using "any" which is an alias for "interface{}"
// to accept any type as a paramenter
func printSomething(data any) {
	/* i:
	   do something based on datatype
	   since we do not have a case for struct
	   and we do not have a defaul case, we will do nothing
	*/
	switch data.(type) {
	case int:
		fmt.Println("This is an integer:", data)
	case float64:
		fmt.Println("This is a float64:", data)
	case string:
		fmt.Println("This is a string:", data)
	}

	// i: alternatives if you want a to do something with a specifig type

	// intVal, ok := data.(int)
	// if ok {
	// 	fmt.Printf("This is the integer + 1: %d\n\n", intVal+1)
	// } else {
	// 	return
	// }
	//
	// floatVal, ok := data.(float64)
	// if ok {
	// 	fmt.Printf("This is the float64 + 1: %f\n\n", floatVal+1.5)
	// } else {
	// 	return
	// }
	//
	// stringVal, ok := data.(string)
	// if ok {
	// 	fmt.Printf("This is the float64 + 1: %s\n\n", stringVal+" world")
	// } else {
	// 	return
	// }
}
