package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func openFile() *os.File {
	f, err := os.Open("input.go.txt")
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func parseLine(line string) (int, int, string, byte) {
	// Split on hifen
	stringList := strings.Split(line, "-")

	pos1, _ := strconv.Atoi(stringList[0])

	// Split on whitespace
	stringList = strings.Split(stringList[1], " ")

	pos2, _ := strconv.Atoi(stringList[0])

	password := stringList[2]

	letter := stringList[1][0]

	return pos1, pos2, password, letter
}

func main() {
	f := openFile()
	defer f.Close()

	valids := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		pos1, pos2, password, letter := parseLine(line)

		matches := 0

		if password[pos1-1] == letter {
			matches++
		}

		if password[pos2-1] == letter {
			matches++
		}

		if matches == 1 {
			valids++
		}
	}

	println("Valids:", valids)
}
