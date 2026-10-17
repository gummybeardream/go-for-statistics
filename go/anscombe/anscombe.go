package anscombe

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func CalculateRegression(data stats.Series) (float64, float64, error) {
	regression, err := stats.LinearRegression(data)
	if err != nil {
		return 0, 0, fmt.Errorf("error calculating linear regression: %w", err)
	}

	slope, intercept := calculateSlope(regression)
	return slope, intercept, nil
}

func calculateSlope(regression stats.Series) (float64, float64) {
	for index, coordinate := range regression {
		if coordinate.X != regression[0].X {
			slope := (regression[index].Y - regression[0].Y) / (regression[index].X - regression[0].X)
			intercept := regression[0].Y - slope*regression[0].X
			return slope, intercept
		}
	}
	return 0, 0
}
