package chapter1

import (
	"math/rand"
	"strings"
)

func shuffle(str string) string {
	if len(str) <= 4 {
		return str
	}
	word := []rune(str)
	for i := len(word) - 2; i >= 1; i-- {
		j := rand.Intn(i) + 1
		word[i], word[j] = word[j], word[i]
	}
	return string(word)
}

func Typoglycemia(str string) string {

	words := strings.Split(str, " ")
	var list []string
	for _, word := range words {
		list = append(list, shuffle(word))
	}
	return strings.Join(list, " ")
}
