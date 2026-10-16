package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func main() {
	datasetOne := stats.Series{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}

	regression, err := stats.LinearRegression(datasetOne)
	if err != nil {
		fmt.Println("Error performing linear regression:", err)
		return
	}

	fmt.Println("Regression:", regression)
	fmt.Println("Dataset 1:", datasetOne)
	fmt.Println("Number of observations:", len(datasetOne))
}
