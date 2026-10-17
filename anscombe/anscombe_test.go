package anscombe_test

import (
	"math"
	"testing"

	"github.com/gummybeardream/go-for-statistics/anscombe"
	"github.com/montanaflynn/stats"
)

func TestCalculateRegression(t *testing.T) {
	tests := []struct {
		name              string
		data              stats.Series
		expectedSlope     float64
		expectedIntercept float64
	}{
		{
			name:              "Dataset I",
			data:              anscombe.DatasetOne(),
			expectedSlope:     0.500091,
			expectedIntercept: 3.000091,
		},
		{
			name:              "Dataset II",
			data:              anscombe.DatasetTwo(),
			expectedSlope:     0.500000,
			expectedIntercept: 3.000909,
		},
		{
			name:              "Dataset III",
			data:              anscombe.DatasetThree(),
			expectedSlope:     0.499727,
			expectedIntercept: 3.002455,
		},
		{
			name:              "Dataset IV",
			data:              anscombe.DatasetFour(),
			expectedSlope:     0.499909,
			expectedIntercept: 3.001727,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test implementation
			slope, intercept, err := anscombe.CalculateRegression(test.data)

			expectedSlope := test.expectedSlope
			expectedIntercept := test.expectedIntercept
			tolerance := 0.00001

			if err != nil {
				t.Fatalf("unexpected regression error: %v", err)
			}
			if math.Abs(slope-expectedSlope) > tolerance {
				t.Errorf("expected %f, got %f", expectedSlope, slope)
			}
			if math.Abs(intercept-expectedIntercept) > tolerance {
				t.Errorf("expected %f, got %f", expectedIntercept, intercept)
			}

		})
	}
}
