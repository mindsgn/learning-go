package main

import "testing"

func TestMakePersonFunctions(t *testing.T) {
	person := MakePerson("Ada", "Lovelace", 36)
	if person.FirstName != "Ada" || person.LastName != "Lovelace" || person.Age != 36 {
		t.Fatalf("unexpected person: %+v", person)
	}
	pointer := MakePersonPointer("Grace", "Hopper", 85)
	if pointer == nil {
		t.Fatal("expected a non-nil pointer")
	}
	if pointer.FirstName != "Grace" || pointer.LastName != "Hopper" || pointer.Age != 85 {
		t.Fatalf("unexpected person pointer: %+v", pointer)
	}
}
