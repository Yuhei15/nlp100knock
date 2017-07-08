package chapter1

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {

	in := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	want := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}
	if got := WordCount(in); !reflect.DeepEqual(got, want) {
		t.Errorf("WordCount(%q) == %q, want %q", in, got, want)
	}
}
