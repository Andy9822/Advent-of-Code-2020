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

		for i := 0; i < EXPECTED_VALUE; i++ {
			complement1 := i
			complement2 := EXPECTED_VALUE - (complement1 + value)

			if complement1 <= 0 || complement2 <= 0 {
				continue
			}

			if vec[value]+vec[complement1]+vec[complement2] == EXPECTED_VALUE {
				fmt.Println(value, "+", complement1, "+", complement2, "=", value+complement1+complement2)
				fmt.Println(value, "*", complement1, "*", complement2, "=", value*complement1*complement2)
				return
			}
		}
	}
}
