package main

import "testing"

func TestColor(t *testing.T) {
	t.Run("color", func(t *testing.T) {
		got := Color("red")
		want := "warm"
		assertCorrectMessage(t, got, want)
	})

	colorTests := []struct {
		color  string
		output string
	}{
		{color: "red", output: "warm"},
		{color: "yellow", output: "warm"},
		{color: "orange", output: "warm"},
		{color: "blue", output: "cool"},
		{color: "green", output: "cool"},
		{color: "purple", output: "cool"},
		{color: "unknown", output: "unknown"},
	}

	for _, tt := range colorTests {
		t.Run(tt.color, func(t *testing.T) {
			got := Color(tt.color)
			want := tt.output
			assertCorrectMessage(t, got, want)
		})
	}
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
