package main

import (
	"math"
	"testing"
)

func TestBuildSquareRootMap(t *testing.T) {
	result := buildSquareRootMap()
	if len(result) != 100000 {
		t.Fatalf("expected 100000 entries, got %d", len(result))
	}
	checks := map[int]float64{0: 0, 1: 1, 81: 9, 99999: math.Sqrt(99999)}
	for key, want := range checks {
		if got := result[key]; math.Abs(got-want) > 1e-9 {
			t.Fatalf("result[%d] = %f, want %f", key, got, want)
		}
	}
}
