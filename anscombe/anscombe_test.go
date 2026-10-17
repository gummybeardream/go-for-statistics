package anscombe_test

import (
	"errors"
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
		expectedErr       error
	}{
		{
			name:              "returns error if dataset is empty",
			data:              stats.Series{},
			expectedSlope:     0,
			expectedIntercept: 0,
			expectedErr:       stats.EmptyInput,
		},
		{
			name:              "Dataset I",
			data:              anscombe.DatasetOne,
			expectedSlope:     0.500091,
			expectedIntercept: 3.000091,
			expectedErr:       nil,
		},
		{
			name:              "Dataset II",
			data:              anscombe.DatasetTwo(),
			expectedSlope:     0.500000,
			expectedIntercept: 3.000909,
			expectedErr:       nil,
		},
		{
			name:              "Dataset III",
			data:              anscombe.DatasetThree(),
			expectedSlope:     0.499727,
			expectedIntercept: 3.002455,
			expectedErr:       nil,
		},
		{
			name:              "Dataset IV",
			data:              anscombe.DatasetFour(),
			expectedSlope:     0.499909,
			expectedIntercept: 3.001727,
			expectedErr:       nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test implementation
			slope, intercept, err := anscombe.CalculateRegression(test.data)

			expectedSlope := test.expectedSlope
			expectedIntercept := test.expectedIntercept
			tolerance := 0.00001

			if !errors.Is(err, test.expectedErr) {
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

func BenchmarkCalculateRegression(b *testing.B) {
	benchmarks := []struct {
		name string
		data stats.Series
	}{
		{
			name: "Dataset I",
			data: anscombe.DatasetOne,
		},
		{
			name: "Dataset II",
			data: anscombe.DatasetTwo(),
		},
		{
			name: "Dataset III",
			data: anscombe.DatasetThree(),
		},
		{
			name: "Dataset IV",
			data: anscombe.DatasetFour(),
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for b.Loop() {
				anscombe.CalculateRegression(benchmark.data)
			}
		})
	}
}
