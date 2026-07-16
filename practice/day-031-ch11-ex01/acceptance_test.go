package main

import (
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestProcessMatchesReference(t *testing.T) {
	for _, args := range [][]string{{"english"}, {"spanish"}, {"french"}, {"unknown"}} {
		dir := testkit.PackageDir(t)
		got := testkit.RunGoRun(t, dir, nil, ".", args...)
		if got.ExitCode != 0 {
			t.Fatalf("student program failed for %v\nstdout:\n%s\nstderr:\n%s", args, got.Stdout, got.Stderr)
		}
		want := testkit.RunGoRun(t, dir, nil, "./reference", args...)
		if want.ExitCode != 0 {
			t.Fatalf("reference program failed for %v\nstdout:\n%s\nstderr:\n%s", args, want.Stdout, want.Stderr)
		}
		if testkit.NormalizeOutput(got.Stdout) != testkit.NormalizeOutput(want.Stdout) {
			t.Fatalf("output mismatch for %v\nwant:\n%s\n\ngot:\n%s", args, testkit.NormalizeOutput(want.Stdout), testkit.NormalizeOutput(got.Stdout))
		}
	}
}
