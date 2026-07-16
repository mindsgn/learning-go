package testkit

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type CmdResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func PackageDir(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	return dir
}

func RunCmd(t *testing.T, dir string, env map[string]string, name string, args ...string) CmdResult {
	t.Helper()
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Env = os.Environ()
	for key, value := range env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	result := CmdResult{
		Stdout: strings.ReplaceAll(stdout.String(), "\r\n", "\n"),
		Stderr: strings.ReplaceAll(stderr.String(), "\r\n", "\n"),
	}
	if err == nil {
		return result
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		result.ExitCode = exitErr.ExitCode()
		return result
	}
	t.Fatalf("run %s %v: %v", name, args, err)
	return CmdResult{}
}

func RunGoRun(t *testing.T, dir string, env map[string]string, pkg string, args ...string) CmdResult {
	t.Helper()
	cmdArgs := []string{"run", pkg}
	cmdArgs = append(cmdArgs, args...)
	return RunCmd(t, dir, env, "go", cmdArgs...)
}

func RunGoTest(t *testing.T, dir string, env map[string]string, args ...string) CmdResult {
	t.Helper()
	cmdArgs := []string{"test"}
	cmdArgs = append(cmdArgs, args...)
	return RunCmd(t, dir, env, "go", cmdArgs...)
}

func NormalizeOutput(in string) string {
	lines := strings.Split(strings.ReplaceAll(in, "\r\n", "\n"), "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimRight(line, " \t")
		if line == "" {
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

func MustReadFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}

func CaptureStdout(t *testing.T, fn func()) string {
	t.Helper()
	oldStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("create stdout pipe: %v", err)
	}
	os.Stdout = writer
	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("close stdout writer: %v", err)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("read captured stdout: %v", err)
	}
	return string(data)
}

func AssertContains(t *testing.T, s string, parts ...string) {
	t.Helper()
	for _, part := range parts {
		if !strings.Contains(s, part) {
			t.Fatalf("expected %q to contain %q", s, part)
		}
	}
}

func AssertMatches(t *testing.T, s string, pattern string) {
	t.Helper()
	re := regexp.MustCompile(pattern)
	if !re.MatchString(s) {
		t.Fatalf("expected %q to match %q", s, pattern)
	}
}

func AssertGoFmtClean(t *testing.T, dir string) {
	t.Helper()
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() && d.Name() == "reference" {
			return filepath.SkipDir
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		formatted, err := format.Source(data)
		if err != nil {
			return fmt.Errorf("%s is not gofmt-clean: %w", path, err)
		}
		if !bytes.Equal(data, formatted) {
			return fmt.Errorf("%s is not gofmt-clean", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func ParseBracketedInts(t *testing.T, raw string) []int {
	t.Helper()
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "[")
	raw = strings.TrimSuffix(raw, "]")
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	parts := strings.Fields(raw)
	out := make([]int, 0, len(parts))
	for _, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			t.Fatalf("parse int %q: %v", part, err)
		}
		out = append(out, n)
	}
	return out
}

func ParseIntLines(t *testing.T, raw string) []int {
	t.Helper()
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	out := make([]int, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			t.Fatalf("parse int line %q: %v", line, err)
		}
		out = append(out, n)
	}
	return out
}
