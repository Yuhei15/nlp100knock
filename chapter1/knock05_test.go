package chapter1

import (
	"reflect"
	"testing"
)

func TestWordNgram(t *testing.T) {
	in := "I am an NLPer"
	want := [][]string{
		{"I", "am"},
		{"am", "an"},
		{"an", "NLPer"},
	}
	if got := WordNgram(in); !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse(%q) == %q, want %q", in, got, want)
	}
}

// [["I" "am"] ["am" "an"] ["an" "NLPer"]],
// want [["I" "am"] ["am" "an"] ["an" "NLPer"]]
func TestCharacterNgram(t *testing.T) {
	in := "I am an NLPer"
	want := [][]string{
		{"a", "m"},
		{"a", "n"},
		{"N", "L"},
		{"L", "P"},
		{"P", "e"},
		{"e", "r"},
	}
	if got := CharacterNgram(in); !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse(%q) == %q, want %q", in, got, want)
	}
}
