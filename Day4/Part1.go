package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func openFile() *os.File {
	f, err := os.Open("input.go.txt")
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func containsAll(passport string, subStrs ...string) bool {
	contains := true
	for _, substring := range subStrs {
		contains = contains && strings.Contains(passport, substring)
	}

	return contains
}

func validatePassport(passport string) bool {
	return containsAll(passport, "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
}

func main() {
	f := openFile()
	defer f.Close()

	passport := ""
	valids := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if validatePassport(passport) {
				valids++
			}
			passport = ""
		} else {
			passport = passport + " " + line
		}
	}

	// Ensure there wasn't a lass passport that ended without blank line but with EOF
	if passport != "" {
		if validatePassport(passport) {
			valids++
		}
	}

	println("Valid passports:", valids)
}
