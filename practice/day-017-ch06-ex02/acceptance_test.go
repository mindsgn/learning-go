package main

import "testing"

func TestUpdateSliceAndGrowSlice(t *testing.T) {
	base := []string{"a", "b", "c"}
	UpdateSlice(base, "done")
	if base[2] != "done" {
		t.Fatalf("UpdateSlice should update the last element, got %v", base)
	}
	originalLen := len(base)
	GrowSlice(base, "extra")
	if len(base) != originalLen {
		t.Fatalf("GrowSlice should not change the caller's slice length, got %d", len(base))
	}
}
