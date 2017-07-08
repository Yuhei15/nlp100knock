package chapter1

import "testing"

func TestTypoglycemia(t *testing.T) {
	in := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	want := "I coundl't bielvee that I cloud actually untdsraend what I was rdnaieg : the pemhnenoal peowr of the huamn mind ."
	if got := Typoglycemia(in); got != want {
		t.Errorf("Typoglycemia(%q) == %q, want %q", in, got, want)
	}
}
