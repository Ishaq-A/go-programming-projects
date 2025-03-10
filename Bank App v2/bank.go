package main

import (
	"fmt"

	"example.com/bank-v2/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------------------------")
		
		panic("Can't continue, sorry.")
	}
	
	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())
	
	for {
		presentOptions()
		
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		
		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			
			if depositAmount <= 0 {
				fmt.Println("Inavlid amount. Must be greater than 0.")
				continue
			}
			
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			
			if withdrawalAmount <= 0 {
				fmt.Println("Inavlid amount. Must be greater than 0.")
				continue
			}
			
			if withdrawalAmount > accountBalance {
				fmt.Println("Inavlid amount. Cannot withdraw more than balance.")
				continue
			}
			
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
		
		fmt.Println("Your choice", choice)
	}
	
}

