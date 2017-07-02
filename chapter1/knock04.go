package chapter1

import "strings"

func ElementSymbol(str string) map[int]string {

	str = strings.Trim(str, ",")
	str = strings.Trim(str, ".")
	words := strings.Split(str, " ")

	list := make(map[int]string)
	for i, word := range words {
		if i == 0 || i == 4 || i == 5 || i == 6 || i == 7 || i == 8 || i == 14 || i == 15 || i == 18 {
			list[i+1] = string([]rune(word)[0])
		} else {
			list[i+1] = string([]rune(word)[0:2])
		}
	}
	return nil
}
