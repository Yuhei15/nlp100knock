package chapter1

import "strings"

func WordCount(str string) []int {

	words := strings.Split(str, " ")

	var list []int
	for _, word := range words {
		word = strings.Trim(word, ",")
		word = strings.Trim(word, ".")
		list = append(list, len(word))
	}
	return list
}
