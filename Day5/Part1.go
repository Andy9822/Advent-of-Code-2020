package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func openFile() *os.File {
	f, err := os.Open("input.go.txt")
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func main() {
	f := openFile()
	defer f.Close()

	maxID := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lowerBound := 0
		upperBound := 128

		for _, char := range line[:7] {
			if char == 'F' {
				upperBound -= (upperBound - lowerBound) / 2
			} else {
				lowerBound += (upperBound - lowerBound) / 2
			}
		}
		row := upperBound - 1

		lowerBound = 0
		upperBound = 8
		for _, char := range line[7:] {
			if char == 'L' {
				upperBound -= (upperBound - lowerBound) / 2
			} else {
				lowerBound += (upperBound - lowerBound) / 2
			}
		}
		seat := upperBound - 1

		if id := row*8 + seat; id > maxID {
			maxID = id
		}
	}

	fmt.Println("Max ID:", maxID)
}
