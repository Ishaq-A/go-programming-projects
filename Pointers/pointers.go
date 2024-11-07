package main

import "fmt"

func main() {
	age := 32 // regular variable
	
	// var agePointer *int = &age
	
	agePointer := &age
	
	fmt.Println("Age:", *agePointer)
	
	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	// dereference pointer
	*age = *age - 18
}

