package main

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello World"
	msg := GetHelloWorld()
	if want != msg {
		t.Fatalf(`exprcted output "Hello World"`)
	}
}
