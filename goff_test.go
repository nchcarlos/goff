package goff

import "testing"

func TestGoff(t *testing.T) {
	want := "Hello, world."
	if got := Goff(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
