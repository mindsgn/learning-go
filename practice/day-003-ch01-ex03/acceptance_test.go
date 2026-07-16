package main

import (
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestProgramIsFormattedAndRuns(t *testing.T) {
	dir := testkit.PackageDir(t)
	testkit.AssertGoFmtClean(t, dir)
	result := testkit.RunGoRun(t, dir, nil, ".")
	if result.ExitCode != 0 {
		t.Fatalf("go run failed\nstdout:\n%s\nstderr:\n%s", result.Stdout, result.Stderr)
	}
	want := "Hello, world!\nPractice makes progress!"
	if testkit.NormalizeOutput(result.Stdout) != want {
		t.Fatalf("unexpected output\nwant:\n%s\n\ngot:\n%s", want, testkit.NormalizeOutput(result.Stdout))
	}
}
