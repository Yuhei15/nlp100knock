package chapter1

func Concatenate(s string) string {

	var rune = []rune(s)
	out := ""
	for i := 0; i < len(rune); i++ {
		if i%2 == 0 {
			out += string(rune[i])
		}
	}
	return string(out)
}
