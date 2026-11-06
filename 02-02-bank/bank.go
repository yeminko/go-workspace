package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.0

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
		default:
			fmt.Println("Invalid choice")
			fmt.Println("Thank you for using Go Bank!")
		}
	}

}
