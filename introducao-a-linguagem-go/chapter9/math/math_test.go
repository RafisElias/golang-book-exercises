package math

import (
	"fmt"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		values []float64
		expect float64
	}{
		{[]float64{1, 2, 3, 4, 5, 6}, 3.5},
		{[]float64{-1, -2, -3, -4, -5, -6}, -3.5},
		{[]float64{-10, 200, 332, 4, -52, -6}, 78},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Retorna a média do slice %v", tt.values)
		t.Run(testName, func(t *testing.T) {
			result := Average(tt.values)
			if result != tt.expect {
				t.Errorf("resultado '%f', esperado '%f'", result, tt.expect)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		values []float64
		expect float64
	}{
		{[]float64{1, 2, 3, 4, 5, 0.1, 6}, 0.1},
		{[]float64{-1, -2, -3, -4, -5, -11, -6}, -11},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Retorna o menor número do slice %v", tt.values)
		t.Run(testName, func(t *testing.T) {
			result := Min(tt.values)
			if result != tt.expect {
				t.Errorf("resultado '%f', esperado '%f'", result, tt.expect)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		values []float64
		expect float64
	}{
		{[]float64{1, 2, 3, 4, 100, 5, 6}, 100},
		{[]float64{-1.3, -2, -3, -1, -5, -11, -6}, -1},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Retorna o maior número do slice %v", tt.values)
		t.Run(testName, func(t *testing.T) {
			result := Max(tt.values)
			if result != tt.expect {
				t.Errorf("resultado '%f', esperado '%f'", result, tt.expect)
			}
		})
	}
}
