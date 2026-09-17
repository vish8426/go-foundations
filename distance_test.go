package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name string
		p, q Point
		want float64
	}{
		{"same point", Point{0, 0}, Point{0, 0}, 0},
		{"unit x", Point{0, 0}, Point{1, 0}, 1},
		{"3-4-5", Point{0, 0}, Point{3, 4}, 5},
		{"negative", Point{-3, -4}, Point{0, 0}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distance(tt.p, tt.q)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
