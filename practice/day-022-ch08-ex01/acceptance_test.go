package main

import "testing"

func TestDoubler(t *testing.T) {
	if got := Doubler(10); got != 20 {
		t.Fatalf("Doubler(10) = %v", got)
	}
	if got := Doubler(11.2); got != 22.4 {
		t.Fatalf("Doubler(11.2) = %v", got)
	}
}
