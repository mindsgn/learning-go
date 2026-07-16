package practice

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindRoot(t *testing.T) {
	root := t.TempDir()
	mustWriteFile(t, filepath.Join(root, "go.mod"), "module learninggo/practice\n")
	mustWriteFile(t, filepath.Join(root, "catalog.json"), "[]\n")
	nested := filepath.Join(root, "day-001", "subdir")
	mustMkdirAll(t, nested)
	got, err := FindRoot(nested)
	if err != nil {
		t.Fatalf("FindRoot returned error: %v", err)
	}
	if got != root {
		t.Fatalf("FindRoot = %s, want %s", got, root)
	}
}

func TestStoreLifecycle(t *testing.T) {
	root := t.TempDir()
	dbPath := filepath.Join(root, "practice.db")
	catalog := []CatalogExercise{
		{Day: 1, Chapter: 1, Exercise: 1, Folder: "day-001-ch01-ex01", ChapterTitle: "Intro", Prompt: "one"},
		{Day: 2, Chapter: 1, Exercise: 2, Folder: "day-002-ch01-ex02", ChapterTitle: "Intro", Prompt: "two"},
	}
	store, err := openWithCatalog(root, dbPath, catalog)
	if err != nil {
		t.Fatalf("openWithCatalog returned error: %v", err)
	}
	defer store.Close()

	summary, err := store.Summary()
	if err != nil {
		t.Fatalf("Summary returned error: %v", err)
	}
	if summary.Total != 2 || summary.Todo != 2 {
		t.Fatalf("unexpected initial summary: %+v", summary)
	}

	started, err := store.Start("1")
	if err != nil {
		t.Fatalf("Start returned error: %v", err)
	}
	if started.Status != "started" {
		t.Fatalf("Start status = %q, want started", started.Status)
	}

	next, err := store.Next()
	if err != nil {
		t.Fatalf("Next returned error: %v", err)
	}
	if next.Day != 1 {
		t.Fatalf("Next day = %d, want 1", next.Day)
	}

	recorded, err := store.RecordTest("1", true)
	if err != nil {
		t.Fatalf("RecordTest returned error: %v", err)
	}
	if !recorded.HasTestResult || !recorded.LastTestPassed || recorded.TestRuns != 1 {
		t.Fatalf("unexpected test result state: %+v", recorded)
	}

	done, err := store.Done("day-001")
	if err != nil {
		t.Fatalf("Done returned error: %v", err)
	}
	if done.Status != "done" {
		t.Fatalf("Done status = %q, want done", done.Status)
	}

	next, err = store.Next()
	if err != nil {
		t.Fatalf("Next after done returned error: %v", err)
	}
	if next.Day != 2 {
		t.Fatalf("Next day after done = %d, want 2", next.Day)
	}
}

func mustWriteFile(t *testing.T, path, contents string) {
	t.Helper()
	mustMkdirAll(t, filepath.Dir(path))
	if err := os.WriteFile(path, []byte(contents), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func mustMkdirAll(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", path, err)
	}
}
