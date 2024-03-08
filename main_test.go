package main

import (
	"testing"
	"time"
)

func TestArgParse(t *testing.T) {
	args := []string{"gosleep"}
	_, err := parseArgs(args)
	if err != ERR_TOO_FEW_ARGS {
		t.Fatalf("Expected Too Few Args Error")
	}

	args = []string{"gosleep", "--test", "fail"}
	_, err = parseArgs(args)
	if err != ERR_TOO_MANY_ARGS {
		t.Fatalf("Expected Too Many Args Error")
	}

	args = []string{"gosleep", "lizard"}
	_, err = parseArgs(args)
	if err == nil {
		t.Fatalf("Expected duration parsing to fail on non-number")
	}

	args = []string{"gosleep", "1"}
	td, err := parseArgs(args)
	if err != nil {
		t.Fatalf("Expected no error parsing int to seconds")
	}
	if td != (1 * time.Second) {
		t.Fatalf("Expected time duration to equal one second")
	}

	args = []string{"gosleep", "1h"}
	td, err = parseArgs(args)
	if err != nil {
		t.Fatalf("Expected no error parsing timestring to duration")
	}
	if td != (1 * time.Hour) {
		t.Fatalf("Expected correct time duration parsing")
	}

	args = []string{"gosleep", "1h30m"}
	td, err = parseArgs(args)
	if err != nil {
		t.Fatalf("Expected no error parsing timestring to duration")
	}
	if td != ((1 * time.Hour) + (30 * time.Minute)) {
		t.Fatalf("Expected correct complex time duration parsing")
	}
}
