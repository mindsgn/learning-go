package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestFuzzWorkflow(t *testing.T) {
	if os.Getenv("PRACTICE_INNER") == "1" {
		t.Skip("inner fuzz run")
	}
	files, err := filepath.Glob("*_test.go")
	if err != nil {
		t.Fatalf("glob *_test.go: %v", err)
	}
	var found bool
	for _, file := range files {
		if file == "acceptance_test.go" {
			continue
		}
		contents := testkit.MustReadFile(t, file)
		if strings.Contains(contents, "func FuzzParseKeyValue") {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("add a fuzz test named FuzzParseKeyValue before running the acceptance test")
	}
	dir := testkit.PackageDir(t)
	unit := testkit.RunGoTest(t, dir, map[string]string{"PRACTICE_INNER": "1"}, "./...")
	if unit.ExitCode != 0 {
		t.Fatalf("inner go test failed\nstdout:\n%s\nstderr:\n%s", unit.Stdout, unit.Stderr)
	}
	fuzz := testkit.RunCmd(t, dir, map[string]string{"PRACTICE_INNER": "1"}, "go", "test", "-run=^$", "-fuzz=FuzzParseKeyValue", "-fuzztime=200ms", "./...")
	if fuzz.ExitCode != 0 {
		t.Fatalf("fuzz run failed\nstdout:\n%s\nstderr:\n%s", fuzz.Stdout, fuzz.Stderr)
	}
}
