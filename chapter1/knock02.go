package chapter1

func MixConcatenate(str1, str2 string) string {

	var rune1, rune2 = []rune(str1), []rune(str2)
	out := ""
	for i := 0; i < len(rune1); i++ {
		out += string(rune1[i : i+1])
		out += string(rune2[i : i+1])
	}
	return string(out)
}
