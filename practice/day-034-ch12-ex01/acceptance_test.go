package main

import (
	"slices"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestProcessDataWritesExpectedValues(t *testing.T) {
	dir := testkit.PackageDir(t)
	result := testkit.RunGoRun(t, dir, nil, ".")
	if result.ExitCode != 0 {
		t.Fatalf("go run failed\nstdout:\n%s\nstderr:\n%s", result.Stdout, result.Stderr)
	}
	got := testkit.ParseIntLines(t, result.Stdout)
	if len(got) != 20 {
		t.Fatalf("expected 20 values, got %d", len(got))
	}
	want := []int{0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 101, 201, 301, 401, 501, 601, 701, 801, 901}
	slices.Sort(got)
	slices.Sort(want)
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("unexpected values: %v", got)
		}
	}
}
