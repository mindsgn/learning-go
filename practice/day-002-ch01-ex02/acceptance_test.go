package main

import (
	"os"
	"path/filepath"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestMakefileCleanRemovesBinary(t *testing.T) {
	dir := testkit.PackageDir(t)
	build := testkit.RunCmd(t, dir, nil, "make", "build")
	if build.ExitCode != 0 {
		t.Fatalf("make build failed\nstdout:\n%s\nstderr:\n%s", build.Stdout, build.Stderr)
	}
	binary := filepath.Join(dir, "hello_world")
	if _, err := os.Stat(binary); err != nil {
		t.Fatalf("expected %s to exist after build: %v", binary, err)
	}
	clean := testkit.RunCmd(t, dir, nil, "make", "clean")
	if clean.ExitCode != 0 {
		t.Fatalf("make clean failed\nstdout:\n%s\nstderr:\n%s", clean.Stdout, clean.Stderr)
	}
	if _, err := os.Stat(binary); !os.IsNotExist(err) {
		t.Fatalf("expected %s to be removed by make clean", binary)
	}
}
