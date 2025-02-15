package main

import (
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	uuid := generate_uuid()

	got := len(uuid)
	want := 36

	// instead of asserting, you basically just raise an error
	// if your test case doesnt match what you expect
	if got != want {
		t.Errorf("want %q got %q", got, want)
	}
}
