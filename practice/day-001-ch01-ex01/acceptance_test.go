package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run program: %v\n%s", err, out)
	}
	got := strings.TrimSpace(string(out))
	if got != "Hello World" {
		t.Errorf("expected %q, got %q", "Hello World", got)
	}
}
