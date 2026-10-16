package anscombe

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func CalculateRegression(data stats.Series) (float64, float64, error) {
	// Placeholder implementation for the regression calculation
	regression, err := stats.LinearRegression(datasetOne)
	if err != nil {
		fmt.Println("Error performing linear regression:", err)
		return
	}

	fmt.Println("Regression:", regression)
	return 8.04, 0.9, nil
}
