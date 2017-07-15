package chapter2

import "testing"

func TestWCLine(t *testing.T) {
	in, want := "hightemp.txt", 24
	if got := WCLine(in); got != want {
		t.Errorf("TestWCLine(%q) == %q, want %q", in, got, want)
	}
}
