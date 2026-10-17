package main

import (
	"fmt"

	"github.com/gummybeardream/go-for-statistics/anscombe"
	"github.com/montanaflynn/stats"
)

func main() {
	datasets := []stats.Series{
		anscombe.DatasetOne(),
		anscombe.DatasetTwo(),
		anscombe.DatasetThree(),
		anscombe.DatasetFour(),
	}

	for index, dataset := range datasets {

		slope, intercept, err := anscombe.CalculateRegression(dataset)
		if err != nil {
			fmt.Println("Regression error:", err)
			continue
		}

		// Display results
		fmt.Printf("Dataset %d: %v\n", index+1, dataset)
		fmt.Println("Number of observations:", len(dataset))
		fmt.Printf("Slope: %5f\n", slope)
		fmt.Printf("Intercept: %5f\n", intercept)

	}

}
