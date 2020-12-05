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

func main() {
	f := openFile()
	defer f.Close()

	trees := 0
	column := -3
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tamanho := len(line)
		column += 3
		if line[column%tamanho] == '#' {
			trees++
		}
	}
	println("\nColision:", trees)
}
