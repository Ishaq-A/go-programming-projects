package main

import "fmt"

// Type Alias
type floatMap map[string]float64

// Custom Method (of the Type being Aliased)
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	
	userNames[0] = "Gohan"
	
	userNames = append(userNames, "Goku")
	userNames = append(userNames, "Vegeta")
	
	fmt.Println(userNames)
	
	courseRatings := make(floatMap, 3)
	
	courseRatings["go"] = 4.7
	courseRatings["react"] = 3.8
	courseRatings["angular"] = 4.7
	
	courseRatings.output()
	
	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}
}

