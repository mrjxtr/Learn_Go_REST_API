package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func newProduct(title string, id int, price float64) Product {
	return Product{
		title: title,
		id:    id,
		price: price,
	}
}

func main() {
	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third array combined as a new list
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.

	// NOTE: #1 SOLUTION
	hobbies := [3]string{"Coding", "Guitar", "Movies"}
	fmt.Println(hobbies)

	// NOTE: #2 SOLUTION
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// NOTE: #3 SOLUTION
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// NOTE: #4 SOLUTION
	extraHobbies := mainHobbies[1:3]
	fmt.Println(extraHobbies)

	// NOTE: #5 SOLUTION
	courseGoals := []string{"Get Good", "Build REST API"}
	fmt.Println(courseGoals)

	// NOTE: #6 SOLUTION
	courseGoals[1] = "Build many projects"
	courseGoals = append(courseGoals, "Get some Money")
	fmt.Println(courseGoals)

	// NOTE: #7 SOLUTION
	// product := []Product{
	// 	{title: "Coffee", id: 1234, price: 149.99},
	// 	{title: "Milk Tea", id: 1212, price: 129.00},
	// }
	product := []Product{
		{"Coffee", 1234, 149.99},
		{"Milk Tea", 1212, 129.00},
	}
	fmt.Println(product)

	matcha := newProduct("Matcha", 1233, 119.99)
	product = append(product, matcha)

	fmt.Println(product)
}
