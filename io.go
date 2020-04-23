package main

import (
	"bufio"
	"log"
	"os"
)

func read(path string) []string {
	currentfile, err := os.Open(path)

	if err != nil {
		log.Fatal("failed to open file: " + path)
	}

	fileScanner := bufio.NewScanner(currentfile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	currentfile.Close()

	return lines
}
