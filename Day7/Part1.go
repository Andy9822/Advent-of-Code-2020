package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var reverseGraph = make(map[string][]string)
var visitedNodes = make(map[string]bool)

func openFile() *os.File {
	f, err := os.Open("input.go.txt")
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func getBagType(phrase string) string {
	regex := regexp.MustCompile(`[0-9]+\s([a-z]+\s[a-z]+)\sbag`)
	return regex.FindStringSubmatch(phrase)[1]
}

func getSonsTypes(str string) []string {
	var sonsTypes []string

	if strings.Contains(str, "no other bags") {
		return sonsTypes
	}

	sons := strings.Split(str, ",")

	for _, son := range sons {
		sonsTypes = append(sonsTypes, getBagType(son))
	}

	return sonsTypes
}

func findAscendentPathsFromNode(node string) int {
	paths := 0
	if !visitedNodes[node] {
		paths += 1
		visitedNodes[node] = true
	}

	for _, parent := range reverseGraph[node] {
		paths += findAscendentPathsFromNode(parent)
	}
	return paths
}

func main() {
	f := openFile()
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := strings.Split(line, "contain")
		rootBag := strings.Split(parsedLine[0], " bags ")[0]
		sonsBags := getSonsTypes(parsedLine[1])

		for _, son := range sonsBags {
			reverseGraph[son] = append(reverseGraph[son], rootBag)
		}
	}

	paths := findAscendentPathsFromNode("shiny gold") - 1

	println("Paths from shiny bag to parent nodes:", paths)
}
