package anscombe

import (
	"errors"

	"github.com/montanaflynn/stats"
)

func CalculateRegression(data stats.Series) (float64, float64, error) {

	// Placeholder implementation for the regression calculation
	regression, err := stats.LinearRegression(data)

	if err != nil {
		return 0, 0, err
	}

	if len(regression) < 2 {
		return 0, 0, errors.New("at least two coordinates are required")
	}

	if regression[1].X == regression[0].X {
		return 0, 0, errors.New("cannot calculate slope: identical X values")

	}

	// Calculate slope
	slope := (regression[1].Y - regression[0].Y) / (regression[1].X - regression[0].X)
	intercept := regression[0].Y - slope*regression[0].X

	// Return placeholder values
	return slope, intercept, nil
}
