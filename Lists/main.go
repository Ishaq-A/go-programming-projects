package main

import "fmt"

func main() {
	// Creating a Slice
	prices := []float64{10.99, 8.99}
	fmt.Println("Prices:", prices)
	
	prices = append(prices, 5.99)
	fmt.Println("Updated Prices:", prices)
	
	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println("Combined Prices:", prices)
}

// EXAMPLE 1 ARRAYS
// func main() {
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	var productNames = [4]string{"A book"}
	
// 	productNames[2] = "A carpet"
	
// 	fmt.Println("Prices:", prices)
// 	fmt.Println("Product Names:", productNames)
// 	fmt.Println(prices[2])
	
// 	featuredPrices := prices[1:3]
// 	highlightedPrice := featuredPrices[:1]
	
// 	fmt.Println("Featured Prices:", featuredPrices)
// 	fmt.Println(highlightedPrice)
	
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }

