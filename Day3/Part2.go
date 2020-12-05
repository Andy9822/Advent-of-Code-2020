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

	columnMove1 := -1
	treescolumnMove1 := 0

	columnMove3 := -3
	treescolumnMove3 := 0

	columnMove5 := -5
	treescolumnMove5 := 0

	columnMove7 := -7
	treescolumnMove7 := 0

	columnMove1Down2 := -1
	treescolumnMove1Down2 := 0

	row := -1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tamanho := len(line)

		columnMove1++
		if line[columnMove1%tamanho] == '#' {
			treescolumnMove1++
		}

		columnMove3 += 3
		if line[columnMove3%tamanho] == '#' {
			treescolumnMove3++
		}

		columnMove5 += 5
		if line[columnMove5%tamanho] == '#' {
			treescolumnMove5++
		}

		columnMove7 += 7
		if line[columnMove7%tamanho] == '#' {
			treescolumnMove7++
		}

		row++
		if row%2 == 0 {
			columnMove1Down2++
			if line[columnMove1Down2%tamanho] == '#' {
				treescolumnMove1Down2++
			}
		}

	}
	println("Colisions on Move right 1, down 1 :", treescolumnMove1)
	println("Colisions on Move right 3, down 1 :", treescolumnMove3)
	println("Colisions on Move right 5, down 1 :", treescolumnMove5)
	println("Colisions on Move right 7, down 1 :", treescolumnMove7)
	println("Colisions on Move right 1, down 2 :", treescolumnMove1Down2)
	println("Mult result: ", treescolumnMove1*treescolumnMove3*treescolumnMove5*treescolumnMove7*treescolumnMove1Down2)
}
