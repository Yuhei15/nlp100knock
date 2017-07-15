package main

import (
	"bufio"
	"fmt"
	"log"
)

import "os"

func wcLine(filePath string) int {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lines
}
