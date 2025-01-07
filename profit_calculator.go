package main

import (
	"errors"
	"fmt"
)

func main() {
	revenue, _ := getUserInput("Revenue: ")
	expenses, _ := getUserInput("Expenses: ")
	taxRate, _ := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getUserInput(text string) (result float64, err error) {
	fmt.Print(text)
	fmt.Scan(&result)

	if result <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
