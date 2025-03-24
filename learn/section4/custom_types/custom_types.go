package main

import "fmt"

type str string // Using custom types/type aliases

func (text *str) log() {
	fmt.Println(*text)
}

func main() {
	var name str = "Jester"

	name.log()
}
