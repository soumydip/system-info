package main

import (
	"errors"
	"testing"
)

func TestParseWMICOutput(t *testing.T) {
	input := "Name\nIntel(R) Core(TM) i7\n\n"
	got := parseWMICOutput(input)
	want := "Intel(R) Core(TM) i7"
	if got != want {
		t.Fatalf("parseWMICOutput() = %q, want %q", got, want)
	}
}

func TestWMICInfoUnavailableOnError(t *testing.T) {
	runner := func(string, ...string) ([]byte, error) {
		return nil, errors.New("wmic not found")
	}

	got := wmicInfo(runner, "cpu", "get", "name")
	if got != "Unavailable" {
		t.Fatalf("wmicInfo() = %q, want %q", got, "Unavailable")
	}
}

func TestFormatBytes(t *testing.T) {
	got := formatBytes("1073741824")
	want := "1.00 GB"
	if got != want {
		t.Fatalf("formatBytes() = %q, want %q", got, want)
	}
}
