package chapter1

import "testing"

func TestConcatenate(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"パタトクカシーー", "パトカー"},
	}
	for _, c := range cases {
		got := Concatenate(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
