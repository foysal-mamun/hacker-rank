package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	mealCostStr, _ := reader.ReadString('\n')
	tipPercentStr, _ := reader.ReadString('\n')
	taxPercentStr, _ := reader.ReadString('\n')

	mealCost, _ := strconv.ParseFloat(strings.TrimSpace(mealCostStr), 64)
	tipPercent, _ := strconv.ParseFloat(strings.TrimSpace(tipPercentStr), 64)
	taxPercent, _ := strconv.ParseFloat(strings.TrimSpace(taxPercentStr), 64)

	tip := mealCost * tipPercent / 100
	tax := mealCost * taxPercent / 100

	totalCost := Round(mealCost + tip + tax)

	fmt.Printf("The total meal cost is %v dollars.\n", totalCost)

}

func Round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}
