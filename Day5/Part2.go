package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var IDs []int

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

		id := row*8 + seat
		IDs = append(IDs, id)
	}

	sort.Ints(IDs)

	lastID := IDs[0]
	missingID := 0

	for _, id := range IDs[1:] {
		if id != lastID+1 {
			missingID = id - 1
			break
		}
		lastID = id
	}
	fmt.Println("Missing ID:", missingID)
}
