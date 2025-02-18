//go:build integration_test

package main

import "testing"

const (
	TEST_VAR = "hello world"
)

func TestDatabase(t *testing.T) {
	want := "hello world"

	if TEST_VAR != want {
		t.Errorf("got %s, want %s", TEST_VAR, want)
	}
}
