package main

import "testing"

func TestRandomSlice(t *testing.T) {
	values := randomSlice(100)
	if len(values) != 100 {
		t.Fatalf("expected 100 values, got %d", len(values))
	}
	for _, value := range values {
		if value < 0 || value >= 100 {
			t.Fatalf("value out of range: %d", value)
		}
	}
	if len(randomSlice(0)) != 0 {
		t.Fatal("randomSlice(0) should return an empty slice")
	}
}
