package main

import (
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestProgramFinishesQuicklyAndPrintsReason(t *testing.T) {
	dir := testkit.PackageDir(t)
	result := testkit.RunGoRun(t, dir, nil, ".")
	if result.ExitCode != 0 {
		t.Fatalf("go run failed\nstdout:\n%s\nstderr:\n%s", result.Stdout, result.Stderr)
	}
	testkit.AssertMatches(t, testkit.NormalizeOutput(result.Stdout), `^total: \d+ number of iterations: \d+ (got 1,234|context deadline exceeded)$`)
}
