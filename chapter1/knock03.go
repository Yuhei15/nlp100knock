package chapter1

import "strings"

func WordCount(str string) []int {

	str = strings.Trim(str, ",")
	str = strings.Trim(str, ".")
	words := strings.Split(str, " ")

	var list []int
	for _, word := range words {
		list = append(list, len(word))
	}
	return list
}
