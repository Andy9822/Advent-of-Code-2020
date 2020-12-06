package main

import (
	"bufio"
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

func countUniqueQuestions(group []string) int {
	lettersMap := make(map[rune]int)
	for _, personQuestions := range group {
		for _, letter := range personQuestions {
			lettersMap[letter] = 1
		}
	}

	return len(lettersMap)
}

func main() {
	f := openFile()
	defer f.Close()

	flags := 0

	var group []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			group = append(group, line)
		} else {
			flags += countUniqueQuestions(group)
			group = nil
		}
	}

	if len(group) != 0 {
		flags += countUniqueQuestions(group)
		group = nil
	}

	println("Flags added:", flags)
}
