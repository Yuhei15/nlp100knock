package chapter1

import "strings"

func WordNgram(str string) [][]string {

	words := strings.Split(str, " ")
	list := [][]string{}
	for i := 0; i < len(words)-1; i++ {
		list = append(list, []string{words[i], words[i+1]})
	}
	return list
}

func CharacterNgram(str string) [][]string {

	words := strings.Split(str, " ")
	list := [][]string{}
	for i := 0; i < len(words); i++ {
		if len(words[i]) >= 2 {
			for j := 0; j < len(words[i])-1; j++ {
				list = append(list, []string{string([]rune(words[i])[j]), string([]rune(words[i])[j+1])})
			}
		}
	}
	return list
}
