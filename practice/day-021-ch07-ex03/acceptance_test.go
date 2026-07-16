package main

import (
	"bytes"
	"testing"
)

type fakeRanker struct{}

func (fakeRanker) Ranking() []string {
	return []string{"France", "Nigeria", "Italy", "India"}
}

func TestRankPrinter(t *testing.T) {
	var out bytes.Buffer
	RankPrinter(fakeRanker{}, &out)
	want := "France\nNigeria\nItaly\nIndia"
	if got := out.String(); got != want {
		t.Fatalf("unexpected output\nwant:\n%s\n\ngot:\n%s", want, got)
	}
}
