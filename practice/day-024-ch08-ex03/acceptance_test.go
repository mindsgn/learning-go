package main

import "testing"

func TestList(t *testing.T) {
	var ints List[int]
	ints.Add(100)
	ints.Add(200)
	ints.Insert(150, 1)
	if got := ints.Index(100); got != 0 {
		t.Fatalf("Index(100) = %d", got)
	}
	if got := ints.Index(150); got != 1 {
		t.Fatalf("Index(150) = %d", got)
	}
	if got := ints.Index(999); got != -1 {
		t.Fatalf("Index(999) = %d", got)
	}
}
