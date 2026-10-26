package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt, profit := calculateEBTAndProfit(revenue, expenses, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Printf("EBT: %v\n", ebt)
	fmt.Printf("Profit: %v\n", profit)
	fmt.Printf("Ratio: %v\n", ratio)

	formattedFV := fmt.Sprintf("EBT: %v\nProfit: %v\nRatio: %v\n", ebt, profit, ratio)
	outputText(formattedFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateEBTAndProfit(revenue, expenses, taxRate float64) (float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	return ebt, profit
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}