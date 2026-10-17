package main

import (
	"fmt"
	"github.com/gummybeardream/go-for-statistics/anscombe"
)

func main() {
	datasetOne := anscombe.DatasetOne()

	fmt.Println("Dataset 1:", datasetOne)
	fmt.Println("Number of observations:", len(datasetOne))

	slope, intercept, err := anscombe.CalculateRegression(datasetOne)
	if err != nil {
		fmt.Println("Regression error:", err)
		return
	}
	fmt.Printf("Slope: %5f\n", slope)
	fmt.Printf("Intercept: %5f\n", intercept)
}
