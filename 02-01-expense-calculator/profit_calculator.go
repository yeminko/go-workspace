package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit := calculateEBTAndProfit(revenue, expenses, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)

	formattedFV := fmt.Sprintf("EBT: %v\nProfit: %v\nRatio: %v\n", ebt, profit, ratio)
	outputText(formattedFV)
}

func getUserInput(infoText string) float64 {
	var input float64
	fmt.Print(infoText)
	fmt.Scan(&input)
	return input
}

func outputText(text string) {
	fmt.Print(text)
}

// Returning Multiple Values - 1
// func calculateEBTAndProfit(revenue, expenses, taxRate float64) (float64, float64) {
// 	ebt := revenue - expenses
// 	profit := ebt * (1 - taxRate/100)
// 	return ebt, profit
// }

// Returning Multiple Values - 2
func calculateEBTAndProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	return
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}