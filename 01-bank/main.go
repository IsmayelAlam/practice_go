package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the bank")

	accountBalance := 100.00

	for {
		printMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println()
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Please try again.")
		}
		fmt.Println()
	}
}
