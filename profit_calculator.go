package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err1 := getUserInput("Revenue: ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	expenses, err2 := getUserInput("Expenses: ")
	if err2 != nil {
		fmt.Println(err1)
		return
	}

	taxRate, err3 := getUserInput("Tax Rate: ")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func getUserInput(text string) (result float64, err error) {
	fmt.Print(text)
	fmt.Scan(&result)

	if result <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
