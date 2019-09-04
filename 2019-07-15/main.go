package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Convert input to int
	income, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// Handle error of string conversion
		fmt.Println(err)
		os.Exit(2)
	}
	taxOwed := tax(income)
	fmt.Printf("Tax Owed: ¤%d", taxOwed)
}

func tax(income int) int {
	// TAX BRACKETS
	//income cap      marginal tax rate
	//  ¤10,000           0.00 (0%)
	//  ¤30,000           0.10 (10%)
	//  ¤100,000          0.25 (25%)
	//    --              0.40 (40%)

	// This function is awful and needs to be changed to not be hardcoded
	// The challenge didn't mind hardcoded values here and I am trying to learn the lang
	// Please, future employer, ignore how awful this function is for readability and adapatability

	taxRateBand1 := 0.10
	taxRateBand2 := 0.25
	taxRateBand3 := 0.40
	// Calculate the max tax from a progressive tax bracket using the size of the bracket
	maxTaxBand1 := 20000 * taxRateBand1 // Difference between 10,000 and 30,000 devided by the tax of the bracket
	maxTaxBand2 := 70000 * taxRateBand2 // Difference between 30,000 and 100,000 devided by the tax of the bracket

	// Take the max tax of prevous bands and calculate what the upper band would be on income. Return the value
	if income > 100000 {
		taxableUpperBand := income - 100000
		tax := float64(taxableUpperBand) * taxRateBand3
		return int(tax + maxTaxBand1 + maxTaxBand2)
	} else if income > 30000 {
		taxableUpperBand := income - 30000
		tax := float64(taxableUpperBand) * taxRateBand2
		return int(tax + maxTaxBand1)
	} else if income > 10000 {
		taxableUpperBand := income - 10000
		tax := float64(taxableUpperBand) * taxRateBand1
		return int(tax)
	} else {
		return income
	}
}
