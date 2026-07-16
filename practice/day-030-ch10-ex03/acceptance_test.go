package adder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestGenericAddAndVersion(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %v", got)
	}
	if got := Add(2.5, 3.5); got != 6.0 {
		t.Fatalf("Add(2.5, 3.5) = %v", got)
	}
	version, err := os.ReadFile("VERSION")
	if err != nil {
		t.Fatalf("reading VERSION: %v", err)
	}
	if strings.TrimSpace(string(version)) != "v2.0.0" {
		t.Fatalf("unexpected version: %q", strings.TrimSpace(string(version)))
	}
	contents := ""
	matches, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatalf("glob *.go: %v", err)
	}
	for _, match := range matches {
		contents += testkit.MustReadFile(t, match)
	}
	testkit.AssertContains(t, contents, "type Number interface", "constraints.Integer", "constraints.Float")
}
