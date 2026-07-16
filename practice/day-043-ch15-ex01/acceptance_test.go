package solver

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestCoverageGoal(t *testing.T) {
	if os.Getenv("PRACTICE_INNER") == "1" {
		t.Skip("inner coverage run")
	}
	files, err := filepath.Glob("*_test.go")
	if err != nil {
		t.Fatalf("glob *_test.go: %v", err)
	}
	userTests := 0
	for _, file := range files {
		if file != "acceptance_test.go" {
			userTests++
		}
	}
	if userTests == 0 {
		t.Fatal("add your own tests for this package before running the acceptance test")
	}
	dir := testkit.PackageDir(t)
	result := testkit.RunGoTest(t, dir, map[string]string{"PRACTICE_INNER": "1"}, "-coverprofile=coverage.out", "./...")
	if result.ExitCode != 0 {
		t.Fatalf("inner go test failed\nstdout:\n%s\nstderr:\n%s", result.Stdout, result.Stderr)
	}
	cover := testkit.RunCmd(t, dir, nil, "go", "tool", "cover", "-func=coverage.out")
	if cover.ExitCode != 0 {
		t.Fatalf("go tool cover failed\nstdout:\n%s\nstderr:\n%s", cover.Stdout, cover.Stderr)
	}
	lines := strings.Split(strings.TrimSpace(cover.Stdout), "\n")
	last := lines[len(lines)-1]
	fields := strings.Fields(last)
	pct := strings.TrimSuffix(fields[len(fields)-1], "%")
	coverage, err := strconv.ParseFloat(pct, 64)
	if err != nil {
		t.Fatalf("parse coverage %q: %v", pct, err)
	}
	if coverage < 90 {
		t.Fatalf("coverage %.1f%% is below the 90%% target", coverage)
	}
}
