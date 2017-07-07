package chapter1

import "strings"

func characterNgram(str string) []string {

	words := strings.Split(str, " ")
	list := []string{}
	for i := 0; i < len(words); i++ {
		if len(words[i]) >= 2 {
			for j := 0; j < len(words[i])-1; j++ {
				list = append(list, string([]rune(words[i])[j])+string([]rune(words[i])[j+1]))
			}
		}
	}
	return list
}

func Contain(str string, slice []string) bool {
	for _, arg := range slice {
		if arg == str {
			return true
		}
	}
	return false
}

func UnionSet(str1 []string, str2 []string) []string {

	set := append(str1, str2...)
	result := []string{}
	encountered := map[string]bool{}
	for _, arg := range set {
		if !encountered[arg] {
			encountered[arg] = true
			result = append(result, arg)
		}
	}
	return result
}

func ProductSet(str1 []string, str2 []string) []string {

	result := []string{}
	encountered := map[string]bool{}
	for _, arg := range str1 {
		if Contain(arg, str2) && !encountered[arg] {
			encountered[arg] = true
			result = append(result, arg)
		}
	}
	return result
}

func DifferenceSet(str1 []string, str2 []string) []string {

	result := []string{}
	encountered := map[string]bool{}
	for _, arg := range str1 {
		if !Contain(arg, str2) && !encountered[arg] {
			encountered[arg] = true
			result = append(result, arg)
		}
	}
	return result
}
