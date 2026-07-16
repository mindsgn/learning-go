package main

import (
	"strings"
	"testing"
)

type validUser struct {
	FirstName string `minStrlen:"3"`
	LastName  string `minStrlen:"4"`
}

type invalidUser struct {
	FirstName string `minStrlen:"5"`
	LastName  string `minStrlen:"6"`
}

func TestValidateStringLength(t *testing.T) {
	if err := ValidateStringLength(validUser{FirstName: "Grace", LastName: "Hopper"}); err != nil {
		t.Fatalf("expected valid user, got %v", err)
	}
	err := ValidateStringLength(invalidUser{FirstName: "Amy", LastName: "Wu"})
	if err == nil {
		t.Fatal("expected an error for short strings")
	}
	message := err.Error()
	if !strings.Contains(message, "FirstName") && !strings.Contains(message, "LastName") {
		t.Fatalf("expected field names in the error, got %q", message)
	}
}
