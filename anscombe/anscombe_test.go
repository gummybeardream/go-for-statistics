package anscombe_test

import (
	"testing"

	"github.com/gummybeardream/go-for-statistics/anscombe"
	"github.com/montanaflynn/stats"
)

func TestCalculateRegression(t *testing.T) {
	t.Run("calculate regression", func(t *testing.T) {
		// Test implementation
		slope, intercept, err := anscombe.CalculateRegression(stats.Series{})

		if slope != .500091 {
			t.Errorf("expected %f, got %f", .500091, slope)
		}
		if intercept != 3.000091 {
			t.Errorf("expected %f, got %f", 3.000091, intercept)
		}
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}
