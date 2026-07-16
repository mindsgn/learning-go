package race

import (
	"os"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestRaceFree(t *testing.T) {
	if os.Getenv("PRACTICE_INNER") == "1" {
		t.Skip("inner race run")
	}
	dir := testkit.PackageDir(t)
	result := testkit.RunGoTest(t, dir, map[string]string{"PRACTICE_INNER": "1"}, "-race", "./...")
	if result.ExitCode != 0 {
		t.Fatalf("go test -race failed\nstdout:\n%s\nstderr:\n%s", result.Stdout, result.Stderr)
	}
}
