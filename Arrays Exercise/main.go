package main

import "fmt"

type Product struct {
	title string
	id string
	price float64
}

func newProduct(prodTitle, prodId string, prodPrice float64) Product {
	return Product{
		title: prodTitle,
		id: prodId,
		price: prodPrice,
	}
}

func main() {
	// Task 1
	hobbies := [3]string{"Anime", "Games", "Cars"}
	fmt.Println(hobbies)
	
	// Task 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	
	// Task 3
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println(slice1, slice2)
	
	// Task 4
	slice3 := slice2[1:3]
	fmt.Println(slice3)
	
	// Task 5
	courseGoals := []string{"GO Expert", "GO Legend"}
	fmt.Println(courseGoals)
	
	// Task 6
	courseGoals[1] = "GO Gangster"
	courseGoals = append(courseGoals, "GO Hero")
	fmt.Println(courseGoals)
	
	// Task 7
	meat := newProduct("Meat", "01", 9.99)
	veg := newProduct("Vegetables", "02", 12.99)
	
	products := []Product{meat, veg}
	fmt.Println(products)
	
	eggs := newProduct("Eggs", "03", 6.99)
	products = append(products, eggs)
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first array that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

