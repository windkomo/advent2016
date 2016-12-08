package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isAba(s string) bool {
	if len(s) != 3 {
		return false
	}

	return s[0] != s[1] && s[0] == s[2]
}

func isBab(bab string, aba string) bool {
	if len(bab) != 3 {
		return false
	}

	return aba[0] == bab[1] && aba[1] == bab[0] && aba[1] == bab[2]
}

func getAbas(s string) []string {
	var abas []string
	for i := 0; i <= len(s)-3; i++ {
		if isAba(s[i : i+3]) {
			abas = append(abas, s[i:i+3])
		}
	}

	return abas
}

func hasBab(s string, aba string) bool {
	for i := 0; i <= len(s)-3; i++ {
		if isBab(s[i:i+3], aba) {
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
			babTest := false
			var abas []string
			rBrackets := regexp.MustCompile("\\[([a-z]+)\\]")
			bracketsStrings := rBrackets.FindAllString(scanner.Text(), -1)

			plainStrings := rBrackets.ReplaceAllString(scanner.Text(), "|")
			plainStringsSlice := strings.Split(plainStrings, "|")

			for i := 0; i < len(plainStringsSlice); i++ {
				abas = append(abas, getAbas(plainStringsSlice[i])...)
			}

			if len(abas) > 0 {
				for i := 0; i < len(bracketsStrings); i++ {
					for _, aba := range abas {
						if hasBab(bracketsStrings[i], aba) {
							babTest = true
							break
						}
					}
				}
			}

			if babTest {
				count++
			}
		}
		fmt.Println(count)
	}
}
