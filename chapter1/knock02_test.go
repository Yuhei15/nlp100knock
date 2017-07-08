package chapter1

import "testing"

func TestMixConcatenate(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want string
	}{
		{"パトカー", "タクシー", "パタトクカシーー"},
	}
	for _, c := range cases {
		if got := MixConcatenate(c.in1, c.in2); got != c.want {
			t.Errorf("MixConcatenate(%q,%q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}
