package main

import "testing"

func TestFileLen(t *testing.T) {
	got, err := fileLen("testdata/sample.txt")
	if err != nil {
		t.Fatalf("fileLen returned unexpected error: %v", err)
	}
	if got != 46 {
		t.Fatalf("fileLen returned %d, want 46", got)
	}
	if _, err := fileLen("testdata/missing.txt"); err == nil {
		t.Fatal("expected an error for a missing file")
	}
}
