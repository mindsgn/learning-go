package main

import "testing"

func TestPrefixer(t *testing.T) {
	hello := prefixer("Hello ")
	if got := hello("Bob"); got != "Hello Bob" {
		t.Fatalf("got %q", got)
	}
	if got := prefixer("Go ")("team"); got != "Go team" {
		t.Fatalf("got %q", got)
	}
}
