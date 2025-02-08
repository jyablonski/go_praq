package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("check counter inc", func(t *testing.T) {

		// this creates a valu e receiver
		// c := Counter{value: 1}

		// this creates a pointer receiver
		c := &Counter{value: 1}

		c.Inc()
		got := c.Value()
		want := 2

		if got != want {
			t.Errorf("expected %q, got %q", got, want)
		}
	})
}
