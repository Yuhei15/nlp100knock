package chapter2

import "bufio"

import "os"

func WCLine(filePath string) int {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	return lines
}
