package adder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestDocsAndVersion(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d", got)
	}
	version, err := os.ReadFile("VERSION")
	if err != nil {
		t.Fatalf("reading VERSION: %v", err)
	}
	if strings.TrimSpace(string(version)) != "v1.0.1" {
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
	testkit.AssertContains(t, contents, "https://www.mathsisfun.com/numbers/addition.html")
	testkit.AssertContains(t, contents, "Package adder")
}
