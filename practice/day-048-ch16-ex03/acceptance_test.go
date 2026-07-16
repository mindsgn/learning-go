package main

import "testing"

import "learninggo/practice/internal/testkit"

func TestMiniCalcViaCgo(t *testing.T) {
	dir := testkit.PackageDir(t)
	cases := []struct {
		args []string
		want string
	}{
		{[]string{"10", "+", "5"}, "15"},
		{[]string{"9", "*", "4"}, "36"},
		{[]string{"12", "-", "3"}, "9"},
		{[]string{"20", "/", "5"}, "4"},
	}
	for _, tc := range cases {
		result := testkit.RunGoRun(t, dir, nil, ".", tc.args...)
		if result.ExitCode != 0 {
			t.Fatalf("go run failed for %v\nstdout:\n%s\nstderr:\n%s", tc.args, result.Stdout, result.Stderr)
		}
		if got := testkit.NormalizeOutput(result.Stdout); got != tc.want {
			t.Fatalf("args %v: got %q, want %q", tc.args, got, tc.want)
		}
	}
}
