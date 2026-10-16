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

		if slope != 8.04 {
			t.Errorf("expected %f, got %f", 8.04, slope)
		}
		if intercept != 0.9 {
			t.Errorf("expected %f, got %f", 0.9, intercept)
		}
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}