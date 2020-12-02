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

	minValue, _ := strconv.Atoi(stringList[0])

	// Split on whitespace
	stringList = strings.Split(stringList[1], " ")

	maxValue, _ := strconv.Atoi(stringList[0])

	password := stringList[2]

	letter := stringList[1][0]

	return minValue, maxValue, password, letter
}

func countLetter(str string, letter byte) int {
	counter := 0
	for i := range str {
		if str[i] == letter {
			counter++
		}
	}

	return counter
}

func main() {
	f := openFile()
	defer f.Close()

	valids := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		minValue, maxValue, password, letter := parseLine(line)

		occurrences := countLetter(password, letter)

		if occurrences >= minValue && occurrences <= maxValue {
			valids++
		}
	}

	println("Valids:", valids)
}
