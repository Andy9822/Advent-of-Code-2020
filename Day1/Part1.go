package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const EXPECTED_VALUE = 2020

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

	var vec [EXPECTED_VALUE]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		vec[value] = value

		if complement := EXPECTED_VALUE - value; vec[value]+vec[complement] == EXPECTED_VALUE {
			fmt.Println(value, "+", complement, "=", value+complement)
			fmt.Println(value, "*", complement, "=", value*complement)
			break
		}
	}
}
