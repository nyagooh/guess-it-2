package main

import (
	"bufio"
	"fmt"
	m "guess/guess"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var previousnumbers []float64
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		previousnumbers = append(previousnumbers, number)
		if len(previousnumbers) > 1 {
			a, b := m.LinearRegression(previousnumbers)
			var nextindex int
			if len(previousnumbers) <= 5 {
				nextindex = len(previousnumbers)
			} else {
				previousnumbers = previousnumbers[len(previousnumbers)-5:]
				nextindex = len(previousnumbers)
			}
			predictvalue := a + b*float64(nextindex)
			margin := 20.0
			lower := math.Round(predictvalue - margin)
			upper := math.Round(predictvalue + margin)
			fmt.Printf("%v %v\n", lower, upper)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

}
