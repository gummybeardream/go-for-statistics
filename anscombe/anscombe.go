package anscombe

import (
	"errors"
	"fmt"

	"github.com/montanaflynn/stats"
)

func CalculateRegression(data stats.Series) (float64, float64, error) {
	regression, err := stats.LinearRegression(data)
	if err != nil {
		return 0, 0, fmt.Errorf("error calculating linear regression: %w", err)
	}

	if len(regression) < 2 {
		return 0, 0, errors.New("at least two coordinates are required")
	}

	slope, intercept, err := calculateSlope(regression)
	if err != nil {
		return 0, 0, fmt.Errorf("error calculating slope: %w", err)
	}
	return slope, intercept, nil
}

func calculateSlope(regression stats.Series) (float64, float64, error) {
	for index, coordinate := range regression {
		if coordinate.X != regression[0].X {
			slope := (regression[index].Y - regression[0].Y) / (regression[index].X - regression[0].X)
			intercept := regression[0].Y - slope*regression[0].X
			return slope, intercept, nil
		}
	}
	return 0, 0, errors.New("identical X values")
}
