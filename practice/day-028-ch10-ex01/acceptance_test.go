package adder

import (
	"os"
	"strings"
	"testing"
)

func TestAddAndVersion(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d", got)
	}
	version, err := os.ReadFile("VERSION")
	if err != nil {
		t.Fatalf("reading VERSION: %v", err)
	}
	if strings.TrimSpace(string(version)) != "v1.0.0" {
		t.Fatalf("unexpected version: %q", strings.TrimSpace(string(version)))
	}
}
