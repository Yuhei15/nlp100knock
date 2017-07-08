package chapter1

import (
	"reflect"
	"testing"
)

func TestElementSymbol(t *testing.T) {
	in := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	want := map[int]string{1: "H", 2: "He", 3: "Li", 4: "Be", 5: "B", 6: "C", 7: "N", 8: "O", 9: "F", 10: "Ne", 11: "Na", 12: "Mi", 13: "Al", 14: "Si", 15: "P", 16: "S", 17: "Cl", 18: "Ar", 19: "K", 20: "Ca"}
	if got := ElementSymbol(in); !reflect.DeepEqual(got, want) {
		t.Errorf("ElementSymbol(%q) == %q, want %q", in, got, want)
	}
}
