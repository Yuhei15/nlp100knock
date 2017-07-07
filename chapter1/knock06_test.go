package chapter1

import (
	"reflect"
	"testing"
)

func TestUnionSet(t *testing.T) {
	in1, in2 := characterNgram("paraparaparadise"), characterNgram("paragraph")
	want := []string{"pa", "ar", "ra", "ap", "ad", "di", "is", "se", "ag", "gr", "ph"}
	if got := UnionSet(in1, in2); !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse(%q,%q) == %q, want %q", in1, in2, got, want)
	}
}

func TestProductSet(t *testing.T) {
	in1, in2 := characterNgram("paraparaparadise"), characterNgram("paragraph")
	want := []string{"pa", "ar", "ra", "ap"}
	if got := ProductSet(in1, in2); !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse(%q,%q) == %q, want %q", in1, in2, got, want)
	}
}

func TestDifferenceSet(t *testing.T) {
	in1, in2 := characterNgram("paraparaparadise"), characterNgram("paragraph")
	want := []string{"ad", "di", "is", "se"}
	if got := DifferenceSet(in1, in2); !reflect.DeepEqual(got, want) {
		t.Errorf("Reverse(%q,%q) == %q, want %q", in1, in2, got, want)
	}
}

func TestContain1(t *testing.T) {
	in, X := "se", characterNgram("paraparaparadise")
	want := true
	if got := Contain(in, X); got != want {
		t.Errorf("Reverse(%q) == %v, want %v", in, got, want)
	}
}

func TestContain2(t *testing.T) {
	in, Y := "se", characterNgram("paragraph")
	want := false
	if got := Contain(in, Y); got != want {
		t.Errorf("Reverse(%q) == %v, want %v", in, got, want)
	}
}
