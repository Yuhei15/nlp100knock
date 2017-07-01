// 文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．

package chapter1

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Printf("%c", s[0])
	return string(runes)
}
