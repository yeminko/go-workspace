package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0.0
	}

	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is: ", accountBalance)
		case 2:
			fmt.Println("Depositing money...")
		case 3:
			fmt.Println("Withdrawing money...")
		case 4:
			fmt.Println("Exiting...")
			writeBalanceToFile(accountBalance)
			return
		default:
			fmt.Println("Invalid choice")
			fmt.Println("Thank you for using Go Bank!")
		}
	}

}
