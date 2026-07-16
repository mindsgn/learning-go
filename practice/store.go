package practice

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

type Progress struct {
	CatalogExercise
	Status         string
	StartedAt      string
	CompletedAt    string
	LastTestAt     string
	LastTestPassed bool
	HasTestResult  bool
	TestRuns       int
}

type Summary struct {
	Total     int
	Todo      int
	Started   int
	Done      int
	LastTaken *Progress
	Next      *Progress
}

type Store struct {
	root string
	db   *sql.DB
}

func FindRoot(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("resolve path %q: %w", start, err)
	}
	for {
		if hasFile(dir, "catalog.json") && hasFile(dir, "go.mod") {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("could not find practice root from %s", start)
}

func DefaultDBPath(root string) string {
	if override := strings.TrimSpace(os.Getenv("PRACTICE_DB")); override != "" {
		return override
	}
	return filepath.Join(root, ".practice.db")
}

func Open(root, dbPath string) (*Store, error) {
	catalog, err := LoadCatalog()
	if err != nil {
		return nil, err
	}
	return openWithCatalog(root, dbPath, catalog)
}

func openWithCatalog(root, dbPath string, catalog []CatalogExercise) (*Store, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database: %w", err)
	}
	store := &Store{root: root, db: db}
	if err := store.init(catalog); err != nil {
		db.Close()
		return nil, err
	}
	return store, nil
}

func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}

func (s *Store) Root() string {
	return s.root
}

func (s *Store) init(catalog []CatalogExercise) error {
	schema := `
CREATE TABLE IF NOT EXISTS exercises (
	day INTEGER PRIMARY KEY,
	chapter INTEGER NOT NULL,
	exercise INTEGER NOT NULL,
	folder TEXT NOT NULL UNIQUE,
	chapter_title TEXT NOT NULL,
	prompt TEXT NOT NULL,
	status TEXT NOT NULL DEFAULT 'todo',
	started_at TEXT NOT NULL DEFAULT '',
	completed_at TEXT NOT NULL DEFAULT '',
	last_test_at TEXT NOT NULL DEFAULT '',
	last_test_passed INTEGER,
	test_runs INTEGER NOT NULL DEFAULT 0,
	updated_at TEXT NOT NULL DEFAULT ''
);`
	if _, err := s.db.Exec(schema); err != nil {
		return fmt.Errorf("create schema: %w", err)
	}
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("begin catalog sync: %w", err)
	}
	stmt, err := tx.Prepare(`
INSERT INTO exercises (day, chapter, exercise, folder, chapter_title, prompt)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT(day) DO UPDATE SET
	chapter = excluded.chapter,
	exercise = excluded.exercise,
	folder = excluded.folder,
	chapter_title = excluded.chapter_title,
	prompt = excluded.prompt`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("prepare catalog sync: %w", err)
	}
	defer stmt.Close()
	for _, exercise := range catalog {
		if _, err := stmt.Exec(
			exercise.Day,
			exercise.Chapter,
			exercise.Exercise,
			exercise.Folder,
			exercise.ChapterTitle,
			exercise.Prompt,
		); err != nil {
			tx.Rollback()
			return fmt.Errorf("sync exercise %d: %w", exercise.Day, err)
		}
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit catalog sync: %w", err)
	}
	return nil
}

func (s *Store) List(filter string) ([]Progress, error) {
	query := `
SELECT day, chapter, exercise, folder, chapter_title, prompt, status, started_at, completed_at, last_test_at, last_test_passed, test_runs
FROM exercises`
	var args []any
	filter = normalizeStatus(filter)
	if filter != "" {
		query += " WHERE status = ?"
		args = append(args, filter)
	}
	query += " ORDER BY day"
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("list exercises: %w", err)
	}
	defer rows.Close()
	return scanProgressRows(rows)
}

func (s *Store) Summary() (Summary, error) {
	var out Summary
	rows, err := s.db.Query(`
SELECT status, COUNT(*)
FROM exercises
GROUP BY status`)
	if err != nil {
		return out, fmt.Errorf("query summary: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return out, fmt.Errorf("scan summary: %w", err)
		}
		out.Total += count
		switch status {
		case "todo":
			out.Todo = count
		case "started":
			out.Started = count
		case "done":
			out.Done = count
		}
	}
	next, err := s.Next()
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return out, err
	}
	if err == nil {
		out.Next = &next
	}
	last, err := s.latestStarted()
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return out, err
	}
	if err == nil {
		out.LastTaken = &last
	}
	return out, nil
}

func (s *Store) Next() (Progress, error) {
	return s.queryOne(`
SELECT day, chapter, exercise, folder, chapter_title, prompt, status, started_at, completed_at, last_test_at, last_test_passed, test_runs
FROM exercises
WHERE status != 'done'
ORDER BY CASE status WHEN 'started' THEN 0 ELSE 1 END, day
LIMIT 1`)
}

func (s *Store) latestStarted() (Progress, error) {
	return s.queryOne(`
SELECT day, chapter, exercise, folder, chapter_title, prompt, status, started_at, completed_at, last_test_at, last_test_passed, test_runs
FROM exercises
WHERE status = 'started'
ORDER BY started_at DESC, day DESC
LIMIT 1`)
}

func (s *Store) Get(selector string) (Progress, error) {
	selector = strings.TrimSpace(selector)
	if selector == "" {
		return Progress{}, fmt.Errorf("missing exercise selector")
	}
	if day, err := strconv.Atoi(selector); err == nil {
		return s.queryOne(`
SELECT day, chapter, exercise, folder, chapter_title, prompt, status, started_at, completed_at, last_test_at, last_test_passed, test_runs
FROM exercises
WHERE day = ?`, day)
	}
	if strings.HasPrefix(selector, "day-") {
		progress, err := s.queryOne(`
SELECT day, chapter, exercise, folder, chapter_title, prompt, status, started_at, completed_at, last_test_at, last_test_passed, test_runs
FROM exercises
WHERE folder = ?`, selector)
		if err == nil {
			return progress, nil
		}
		if !errors.Is(err, sql.ErrNoRows) {
			return Progress{}, err
		}
		progresses, err := s.queryMany(`
SELECT day, chapter, exercise, folder, chapter_title, prompt, status, started_at, completed_at, last_test_at, last_test_passed, test_runs
FROM exercises
WHERE folder LIKE ?
ORDER BY day`, selector+"%")
		if err != nil {
			return Progress{}, err
		}
		switch len(progresses) {
		case 0:
			return Progress{}, sql.ErrNoRows
		case 1:
			return progresses[0], nil
		default:
			return Progress{}, fmt.Errorf("selector %q matches multiple exercises; use the full folder name", selector)
		}
	}
	return Progress{}, fmt.Errorf("could not understand selector %q; use a day number like 4 or a folder like day-004-ch02-ex01", selector)
}

func (s *Store) Start(selector string) (Progress, error) {
	progress, err := s.Get(selector)
	if err != nil {
		return Progress{}, err
	}
	now := time.Now().UTC().Format(time.RFC3339)
	_, err = s.db.Exec(`
UPDATE exercises
SET status = CASE WHEN status = 'done' THEN 'done' ELSE 'started' END,
	started_at = CASE WHEN started_at = '' THEN ? ELSE started_at END,
	updated_at = ?
WHERE day = ?`, now, now, progress.Day)
	if err != nil {
		return Progress{}, fmt.Errorf("mark start: %w", err)
	}
	return s.Get(strconv.Itoa(progress.Day))
}

func (s *Store) Done(selector string) (Progress, error) {
	progress, err := s.Get(selector)
	if err != nil {
		return Progress{}, err
	}
	now := time.Now().UTC().Format(time.RFC3339)
	_, err = s.db.Exec(`
UPDATE exercises
SET status = 'done',
	started_at = CASE WHEN started_at = '' THEN ? ELSE started_at END,
	completed_at = ?,
	updated_at = ?
WHERE day = ?`, now, now, now, progress.Day)
	if err != nil {
		return Progress{}, fmt.Errorf("mark done: %w", err)
	}
	return s.Get(strconv.Itoa(progress.Day))
}

func (s *Store) RecordTest(selector string, passed bool) (Progress, error) {
	progress, err := s.Get(selector)
	if err != nil {
		return Progress{}, err
	}
	now := time.Now().UTC().Format(time.RFC3339)
	_, err = s.db.Exec(`
UPDATE exercises
SET last_test_at = ?,
	last_test_passed = ?,
	test_runs = test_runs + 1,
	updated_at = ?
WHERE day = ?`, now, boolToInt(passed), now, progress.Day)
	if err != nil {
		return Progress{}, fmt.Errorf("record test result: %w", err)
	}
	return s.Get(strconv.Itoa(progress.Day))
}

func (s *Store) ExercisePath(progress Progress) string {
	return filepath.Join(s.root, progress.Folder)
}

func (s *Store) queryOne(query string, args ...any) (Progress, error) {
	row := s.db.QueryRow(query, args...)
	return scanProgressRow(row)
}

func (s *Store) queryMany(query string, args ...any) ([]Progress, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("query exercises: %w", err)
	}
	defer rows.Close()
	return scanProgressRows(rows)
}

func scanProgressRows(rows *sql.Rows) ([]Progress, error) {
	var out []Progress
	for rows.Next() {
		progress, err := scanProgress(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, progress)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate exercises: %w", err)
	}
	return out, nil
}

func scanProgressRow(row *sql.Row) (Progress, error) {
	return scanProgress(row)
}

type scanner interface {
	Scan(dest ...any) error
}

func scanProgress(s scanner) (Progress, error) {
	var progress Progress
	var testPassed sql.NullInt64
	err := s.Scan(
		&progress.Day,
		&progress.Chapter,
		&progress.Exercise,
		&progress.Folder,
		&progress.ChapterTitle,
		&progress.Prompt,
		&progress.Status,
		&progress.StartedAt,
		&progress.CompletedAt,
		&progress.LastTestAt,
		&testPassed,
		&progress.TestRuns,
	)
	if err != nil {
		return Progress{}, err
	}
	if testPassed.Valid {
		progress.HasTestResult = true
		progress.LastTestPassed = testPassed.Int64 == 1
	}
	return progress, nil
}

func normalizeStatus(status string) string {
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "", "all":
		return ""
	case "todo":
		return "todo"
	case "started", "in-progress", "inprogress":
		return "started"
	case "done", "complete", "completed":
		return "done"
	default:
		return status
	}
}

func hasFile(dir, name string) bool {
	info, err := os.Stat(filepath.Join(dir, name))
	return err == nil && !info.IsDir()
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
