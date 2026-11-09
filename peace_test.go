package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunFile(t *testing.T) {
	var output bytes.Buffer
	if err := NewInterpreter().RunFile("hello.peace", &output); err != nil {
		t.Fatal(err)
	}
	if got, want := output.String(), "10\n15\n"; got != want {
		t.Fatalf("RunFile output = %q, want %q", got, want)
	}
}

func TestRunLineReportsUndefinedVariables(t *testing.T) {
	err := NewInterpreter().RunLine("print missing", &bytes.Buffer{})
	if err == nil || !strings.Contains(err.Error(), "undefined variable: missing") {
		t.Fatalf("RunLine error = %v, want undefined-variable error", err)
	}
}

func TestRunLineSupportsLargeIntegersAndTypeAnnotation(t *testing.T) {
	var output bytes.Buffer
	interpreter := NewInterpreter()
	if err := interpreter.RunLine("set x = 999999999999999999999999 : type integer", &output); err != nil {
		t.Fatal(err)
	}
	if err := interpreter.RunLine("print (+ x 1)", &output); err != nil {
		t.Fatal(err)
	}
	if got, want := output.String(), "1000000000000000000000000\n"; got != want {
		t.Fatalf("RunLine output = %q, want %q", got, want)
	}
}

func TestRunREPLEvaluatesEnteredStatements(t *testing.T) {
	var output, errorOutput bytes.Buffer
	input := strings.NewReader("set x = 7\nprint (+ x 5)\n:quit\n")
	if err := NewInterpreter().RunREPL(input, &output, &errorOutput); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "12\n") {
		t.Fatalf("REPL output = %q, want it to contain result 12", output.String())
	}
	if errorOutput.Len() != 0 {
		t.Fatalf("REPL error output = %q, want empty", errorOutput.String())
	}
}

func TestRunLineRejectsUnsupportedTypeAnnotation(t *testing.T) {
	err := NewInterpreter().RunLine("set x = 42 : type string", &bytes.Buffer{})
	if err == nil || !strings.Contains(err.Error(), "unsupported type: string") {
		t.Fatalf("RunLine error = %v, want unsupported type error", err)
	}
}

func TestRunLineUsesGlobalEnvironmentAcrossStatements(t *testing.T) {
	var output bytes.Buffer
	interpreter := NewInterpreter()
	if err := interpreter.RunLine("set x = 10", &output); err != nil {
		t.Fatal(err)
	}
	if err := interpreter.RunLine("print (+ x 5)", &output); err != nil {
		t.Fatal(err)
	}
	if got, want := output.String(), "15\n"; got != want {
		t.Fatalf("RunLine output = %q, want %q", got, want)
	}
}

func TestRunLineSupportsReassignmentAndArithmeticOperators(t *testing.T) {
	var output bytes.Buffer
	interpreter := NewInterpreter()
	if err := interpreter.RunLine("set x = 10", &output); err != nil {
		t.Fatal(err)
	}
	if err := interpreter.RunLine("set x = 20", &output); err != nil {
		t.Fatal(err)
	}
	if err := interpreter.RunLine("print (- x 3)", &output); err != nil {
		t.Fatal(err)
	}
	if err := interpreter.RunLine("print (* x 2)", &output); err != nil {
		t.Fatal(err)
	}
	if err := interpreter.RunLine("print (/ x 4)", &output); err != nil {
		t.Fatal(err)
	}
	if got, want := output.String(), "17\n40\n5\n"; got != want {
		t.Fatalf("RunLine output = %q, want %q", got, want)
	}
}
