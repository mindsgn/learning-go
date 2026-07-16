package main

import (
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestToolingWorkflow(t *testing.T) {
	if placeholder := strings.TrimSpace(testkit.MustReadFile(t, filepath.Join(testkit.PackageDir(t), "exercise_stub.go"))); placeholder == "" {
		// keep the package non-empty even before the user starts.
	}
	if strings.TrimSpace(strings.Join(userTestFiles(t), "")) == "" {
		t.Fatal("add at least one user-written _test.go file in this folder")
	}
	if innerCoverage := innerGoTestCoverage(t); innerCoverage < 60 {
		t.Fatalf("coverage %.1f%% is below the 60%% target", innerCoverage)
	}
	vet := testkit.RunCmd(t, testkit.PackageDir(t), nil, "go", "vet", "./...")
	if vet.ExitCode != 0 {
		t.Fatalf("go vet failed\nstdout:\n%s\nstderr:\n%s", vet.Stdout, vet.Stderr)
	}
}

func userTestFiles(t *testing.T) []string {
	files, err := filepath.Glob("*_test.go")
	if err != nil {
		t.Fatalf("glob *_test.go: %v", err)
	}
	out := make([]string, 0, len(files))
	for _, file := range files {
		if file == "acceptance_test.go" {
			continue
		}
		out = append(out, file)
	}
	return out
}

func innerGoTestCoverage(t *testing.T) float64 {
	t.Helper()
	if strings.TrimSpace(strings.Join(userTestFiles(t), "")) == "" {
		return 0
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
	value, err := strconv.ParseFloat(pct, 64)
	if err != nil {
		t.Fatalf("parse coverage %q: %v", pct, err)
	}
	return value
}
