package chapter1

import "testing"

func TestCipher(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"I'm a perfect human", "I'n z kviuvxg sfnzm"},
		{"I'n z kviuvxg sfnzm", "I'm a perfect human"},
	}
	for _, c := range cases {
		if got := Cipher(c.in); got != c.want {
			t.Errorf("Cipher(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
