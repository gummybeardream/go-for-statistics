package main

import (
	"fmt"

	"github.com/gummybeardream/go-for-statistics/anscombe"
	"github.com/montanaflynn/stats"
)

func main() {
	datasets := []stats.Series{
		anscombe.DatasetOne,
		anscombe.DatasetTwo,
		anscombe.DatasetThree,
		anscombe.DatasetFour,
	}

	for index, dataset := range datasets {
		slope, intercept, err := anscombe.CalculateRegression(dataset)
		if err != nil {
			fmt.Printf("failed to calculate regression for dataset %d, received error: %s", index+1, err.Error())
			continue
		}
		regressionResults(index, slope, intercept, dataset)
	}
}

func regressionResults(i int, slope, intercept float64, dataset stats.Series) {
	fmt.Printf("Dataset %d: %v\n", i+1, dataset)
	fmt.Println("Number of observations:", len(dataset))
	fmt.Printf("Slope: %5f\n", slope)
	fmt.Printf("Intercept: %5f\n", intercept)
}
