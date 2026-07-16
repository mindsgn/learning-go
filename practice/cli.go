package practice

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

func RunCLI(stdout, stderr io.Writer, args []string) int {
	root, err := currentRoot()
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	store, err := Open(root, DefaultDBPath(root))
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	defer store.Close()

	if len(args) == 0 {
		printHelp(stdout)
		return 0
	}

	switch args[0] {
	case "help", "-h", "--help":
		printHelp(stdout)
		return 0
	case "list":
		filter := ""
		if len(args) > 1 {
			filter = args[1]
		}
		return runList(stdout, stderr, store, filter)
	case "status":
		return runStatus(stdout, stderr, store)
	case "next":
		return runNext(stdout, stderr, store)
	case "show":
		if len(args) < 2 {
			fmt.Fprintln(stderr, "usage: practice show <day|folder>")
			return 1
		}
		return runShow(stdout, stderr, store, args[1])
	case "start":
		if len(args) < 2 {
			fmt.Fprintln(stderr, "usage: practice start <day|folder>")
			return 1
		}
		return runStart(stdout, stderr, store, args[1])
	case "done":
		if len(args) < 2 {
			fmt.Fprintln(stderr, "usage: practice done <day|folder>")
			return 1
		}
		return runDone(stdout, stderr, store, args[1])
	case "test":
		selector := ""
		extra := []string{"test"}
		if len(args) > 1 {
			selector = args[1]
			if len(args) > 2 {
				if args[2] == "--" {
					extra = append(extra, args[3:]...)
				} else {
					extra = append(extra, args[2:]...)
				}
			}
		}
		return runTest(stdout, stderr, store, selector, extra)
	default:
		fmt.Fprintf(stderr, "unknown command %q\n\n", args[0])
		printHelp(stderr)
		return 1
	}
}

func runList(stdout, stderr io.Writer, store *Store, filter string) int {
	progresses, err := store.List(filter)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	tw := tabwriter.NewWriter(stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(tw, "DAY\tSTATUS\tTESTS\tLAST TEST\tFOLDER\tCHAPTER")
	for _, progress := range progresses {
		last := "-"
		if progress.HasTestResult {
			if progress.LastTestPassed {
				last = "pass"
			} else {
				last = "fail"
			}
		}
		fmt.Fprintf(
			tw,
			"%03d\t%s\t%d\t%s\t%s\tChapter %d: %s\n",
			progress.Day,
			progress.Status,
			progress.TestRuns,
			last,
			progress.Folder,
			progress.Chapter,
			progress.ChapterTitle,
		)
	}
	tw.Flush()
	return 0
}

func runStatus(stdout, stderr io.Writer, store *Store) int {
	summary, err := store.Summary()
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "Total: %d\n", summary.Total)
	fmt.Fprintf(stdout, "Todo: %d\n", summary.Todo)
	fmt.Fprintf(stdout, "Started: %d\n", summary.Started)
	fmt.Fprintf(stdout, "Done: %d\n", summary.Done)
	if summary.LastTaken != nil {
		fmt.Fprintf(stdout, "\nCurrent: day %03d (%s)\n", summary.LastTaken.Day, summary.LastTaken.Folder)
	}
	if summary.Next != nil {
		fmt.Fprintf(stdout, "Next: day %03d (%s)\n", summary.Next.Day, summary.Next.Folder)
	} else {
		fmt.Fprintln(stdout, "Next: all exercises are marked done")
	}
	return 0
}

func runNext(stdout, stderr io.Writer, store *Store) int {
	progress, err := store.Next()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Fprintln(stdout, "All exercises are marked done.")
			return 0
		}
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	printExercise(stdout, store, progress)
	return 0
}

func runShow(stdout, stderr io.Writer, store *Store, selector string) int {
	progress, err := store.Get(selector)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Fprintf(stderr, "exercise %q was not found\n", selector)
			return 1
		}
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	printExercise(stdout, store, progress)
	return 0
}

func runStart(stdout, stderr io.Writer, store *Store, selector string) int {
	progress, err := store.Start(selector)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "Started day %03d: %s\n", progress.Day, progress.Folder)
	fmt.Fprintf(stdout, "Path: %s\n", store.ExercisePath(progress))
	return 0
}

func runDone(stdout, stderr io.Writer, store *Store, selector string) int {
	progress, err := store.Done(selector)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	fmt.Fprintf(stdout, "Marked day %03d done: %s\n", progress.Day, progress.Folder)
	return 0
}

func runTest(stdout, stderr io.Writer, store *Store, selector string, args []string) int {
	var progress Progress
	var err error
	if strings.TrimSpace(selector) == "" {
		progress, err = store.Next()
	} else {
		progress, err = store.Get(selector)
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Fprintln(stderr, "no matching exercise found")
			return 1
		}
		fmt.Fprintf(stderr, "error: %v\n", err)
		return 1
	}
	cmd := exec.Command("go", args...)
	cmd.Dir = store.ExercisePath(progress)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	passed := err == nil
	if _, recordErr := store.RecordTest(fmt.Sprintf("%d", progress.Day), passed); recordErr != nil {
		fmt.Fprintf(stderr, "warning: could not record test result: %v\n", recordErr)
	}
	if err == nil {
		fmt.Fprintf(stdout, "\nRecorded passing test for day %03d.\n", progress.Day)
		return 0
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		fmt.Fprintf(stderr, "\nRecorded failing test for day %03d.\n", progress.Day)
		return exitErr.ExitCode()
	}
	fmt.Fprintf(stderr, "error: %v\n", err)
	return 1
}

func printHelp(w io.Writer) {
	fmt.Fprintln(w, "practice <command>")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Commands:")
	fmt.Fprintln(w, "  list [todo|started|done]   List exercises and progress")
	fmt.Fprintln(w, "  status                     Show progress summary")
	fmt.Fprintln(w, "  next                       Show the next exercise to work on")
	fmt.Fprintln(w, "  show <day|folder>          Show one exercise in detail")
	fmt.Fprintln(w, "  start <day|folder>         Mark an exercise as started")
	fmt.Fprintln(w, "  done <day|folder>          Mark an exercise as done")
	fmt.Fprintln(w, "  test [day|folder] [-- ...] Run go test in an exercise folder and record the result")
}

func printExercise(w io.Writer, store *Store, progress Progress) {
	fmt.Fprintf(w, "Day: %03d\n", progress.Day)
	fmt.Fprintf(w, "Chapter: %d - %s\n", progress.Chapter, progress.ChapterTitle)
	fmt.Fprintf(w, "Exercise: %d\n", progress.Exercise)
	fmt.Fprintf(w, "Status: %s\n", progress.Status)
	fmt.Fprintf(w, "Folder: %s\n", progress.Folder)
	fmt.Fprintf(w, "Path: %s\n", store.ExercisePath(progress))
	if progress.StartedAt != "" {
		fmt.Fprintf(w, "Started: %s\n", progress.StartedAt)
	}
	if progress.CompletedAt != "" {
		fmt.Fprintf(w, "Completed: %s\n", progress.CompletedAt)
	}
	if progress.LastTestAt != "" {
		result := "fail"
		if progress.LastTestPassed {
			result = "pass"
		}
		fmt.Fprintf(w, "Last Test: %s (%s)\n", progress.LastTestAt, result)
	}
	fmt.Fprintf(w, "Test Runs: %d\n", progress.TestRuns)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Prompt:")
	fmt.Fprintln(w, progress.Prompt)
}

func currentRoot() (string, error) {
	if override := strings.TrimSpace(os.Getenv("PRACTICE_ROOT")); override != "" {
		return filepath.Abs(override)
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}
	return FindRoot(wd)
}
