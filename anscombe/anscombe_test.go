package anscombe_test

import (
	"math"
	"testing"

	"github.com/gummybeardream/go-for-statistics/anscombe"
)

func TestCalculateRegression(t *testing.T) {
	testData := anscombe.DatasetOne()

	t.Run("calculate regression", func(t *testing.T) {
		// Test implementation
		slope, intercept, err := anscombe.CalculateRegression(testData)

		expectedSlope := 0.500091
		expectedIntercept := 3.000091
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
