package main

import (
	"path/filepath"
	"strings"
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestOfflineShareNoteExists(t *testing.T) {
	dir := testkit.PackageDir(t)
	data := testkit.MustReadFile(t, filepath.Join(dir, "playground_share.md"))
	trimmed := strings.TrimSpace(data)
	if trimmed == "" {
		t.Fatal("playground_share.md is empty")
	}
	if !strings.Contains(data, "Hello, world") && !strings.Contains(data, "go.dev/play") && !strings.Contains(data, "play.golang.org") {
		t.Fatal("expected a playground link or a local Hello World note")
	}
}
