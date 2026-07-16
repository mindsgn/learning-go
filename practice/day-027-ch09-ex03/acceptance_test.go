package main

import (
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestMatchesReference(t *testing.T) {
	dir := testkit.PackageDir(t)
	got := testkit.RunGoRun(t, dir, nil, ".")
	if got.ExitCode != 0 {
		t.Fatalf("student program failed\nstdout:\n%s\nstderr:\n%s", got.Stdout, got.Stderr)
	}
	want := testkit.RunGoRun(t, dir, nil, "./reference")
	if want.ExitCode != 0 {
		t.Fatalf("reference program failed\nstdout:\n%s\nstderr:\n%s", want.Stdout, want.Stderr)
	}
	if testkit.NormalizeOutput(got.Stdout) != testkit.NormalizeOutput(want.Stdout) {
		t.Fatalf("output mismatch\n\nwant:\n%s\n\ngot:\n%s", testkit.NormalizeOutput(want.Stdout), testkit.NormalizeOutput(got.Stdout))
	}
}
