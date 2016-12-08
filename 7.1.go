package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isAbba(s string) bool {
	if len(s) != 4 {
		return false
	}

	return s[0] != s[1] && s[0] == s[3] && s[1] == s[2]
}

func hasAbba(s string) bool {
	for i := 0; i <= len(s)-4; i++ {
		if isAbba(s[i : i+4]) {
			return true
		}
	}

	return false
}

func main() {

	if file, err := os.Open("input.7.txt"); err == nil {

		// make sure it gets closed
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0

		for scanner.Scan() {
			bracketsTest := true
			plainTest := false
			rBrackets := regexp.MustCompile("\\[([a-z]+)\\]")
			bracketsStrings := rBrackets.FindAllString(scanner.Text(), -1)

			for i := 0; i < len(bracketsStrings); i++ {
				if hasAbba(bracketsStrings[i]) {
					bracketsTest = false
          if hasAbba(bracketsStrings[i]) {
            fmt.Println(bracketsStrings[i])
          }
					break
				}
			}

			plainStrings := rBrackets.ReplaceAllString(scanner.Text(), "|")
			plainStringsSlice := strings.Split(plainStrings, "|")

			for i := 0; i < len(plainStringsSlice); i++ {
				if hasAbba(plainStringsSlice[i]) {
					plainTest = true
					break
				}
			}

			if plainTest && bracketsTest {
				count++
			}
		}
		fmt.Println(count)
	}
}
