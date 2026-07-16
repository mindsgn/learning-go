package main

import (
	"reflect"
	"testing"
)

func TestLeagueMethods(t *testing.T) {
	league := League{
		Name: "Test League",
		Teams: map[string]Team{
			"A": {Name: "A", Players: []string{"p1"}},
			"B": {Name: "B", Players: []string{"p1"}},
			"C": {Name: "C", Players: []string{"p1"}},
		},
		Wins: map[string]int{},
	}
	league.MatchResult("A", 90, "B", 80)
	league.MatchResult("A", 85, "C", 70)
	league.MatchResult("B", 88, "C", 81)
	got := league.Ranking()
	want := []string{"A", "B", "C"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Ranking() = %v, want %v", got, want)
	}
}
