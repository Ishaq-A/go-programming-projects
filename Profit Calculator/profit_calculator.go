package main

import "fmt"

func main() {
	var taxRate float64
	var revenue float64
	var expenses float64
	
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)
	
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

