package chapter1

import "regexp"

func Cipher(clearText string) string {

	cipherText := ""
	for _, value := range clearText {
		if regexp.MustCompile(`[a-z]`).Match([]byte(string(value))) {
			cipherText += string(219 - value)
		} else {
			cipherText += string(value)
		}
	}
	return cipherText
}
