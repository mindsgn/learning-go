package main

import "testing"

func TestBuildPeople(t *testing.T) {
	people := buildPeople(5)
	if len(people) != 5 {
		t.Fatalf("expected 5 people, got %d", len(people))
	}
	for _, person := range people {
		if person.FirstName == "" || person.LastName == "" {
			t.Fatalf("expected populated people, got %+v", person)
		}
	}
	if len(buildPeople(0)) != 0 {
		t.Fatal("buildPeople(0) should return an empty slice")
	}
}
