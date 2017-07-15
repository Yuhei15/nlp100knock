package chapter2

import "bufio"

import "os"

func WCLine(filePath string) int {

	file, _ := os.Open(filePath)
	defer file.Close()

	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}
	return lines
}
