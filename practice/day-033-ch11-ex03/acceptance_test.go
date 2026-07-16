package main

import (
	"path/filepath"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestCrossCompile(t *testing.T) {
	dir := testkit.PackageDir(t)
	result := testkit.RunCmd(t, dir, map[string]string{"GOOS": "windows", "GOARCH": "arm64"}, "go", "build", "-o", filepath.Join("dist", "rights.exe"), ".")
	if result.ExitCode != 0 {
		t.Fatalf("cross-compile failed\nstdout:\n%s\nstderr:\n%s", result.Stdout, result.Stderr)
	}
}
