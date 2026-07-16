package main

import "testing"

func TestDescribeNumber(t *testing.T) {
	tests := map[int]string{
		2:  "Two!",
		3:  "Three!",
		6:  "Six!",
		7:  "Never mind",
		12: "Six!",
	}
	for input, want := range tests {
		if got := describeNumber(input); got != want {
			t.Fatalf("describeNumber(%d) = %q, want %q", input, got, want)
		}
	}
}
