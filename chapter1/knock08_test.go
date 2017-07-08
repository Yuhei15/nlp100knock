package chapter1

import "testing"

func TestStringsJoin(t *testing.T) {
	in1, in2, in3 := "12", "気温", "22.4"
	want := "12の時気温は22.4"
	if got := StringsJoin(in1, in2, in3); got != want {
		t.Errorf("StringsJoin(%q,%q,%q) == %q, want %q", in1, in2, in3, got, want)
	}
}
