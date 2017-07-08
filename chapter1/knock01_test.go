package chapter1

import "testing"

func TestConcatenate(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"パタトクカシーー", "パトカー"},
	}
	for _, c := range cases {
		if got := Concatenate(c.in); got != c.want {
			t.Errorf("Concatenate(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
