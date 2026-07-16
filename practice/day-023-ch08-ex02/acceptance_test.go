package main

import (
	"testing"

	"learninggo/practice/internal/testkit"
)

func TestPrintIt(t *testing.T) {
	output := testkit.CaptureStdout(t, func() {
		PrintIt(PrintInt(20))
		PrintIt(PrintFloat(10.23))
	})
	want := "20\n10.230000"
	if got := testkit.NormalizeOutput(output); got != want {
		t.Fatalf("unexpected output\nwant:\n%s\n\ngot:\n%s", want, got)
	}
}
