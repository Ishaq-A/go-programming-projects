package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	
	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	
	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	
	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)
	
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
	
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	
	// Outputs Information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (Real/Infalted):", futureRealValue)
	
	// Outputs Formatted String
	// fmt.Printf("Future Value: %.1f\n", futureValue)
	// fmt.Printf("Future Value (Real/Inflated): %.1f\n", futureRealValue)
	
	// Creating Formatted String
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Real/Inflated): %.1f\n", futureRealValue)
	
	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	
	return
}

