package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	"regexp"
	"strconv"
)

func main() {

	if file, err := os.Open("input.3.txt"); err == nil {

		// make sure it gets closed
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0

		for scanner.Scan() {
			re := regexp.MustCompile("[0-9]+")
			lengths := re.FindAllString(scanner.Text(), -1)

			length1, _ := strconv.Atoi(lengths[0])
			length2, _ := strconv.Atoi(lengths[1])
			length3, _ := strconv.Atoi(lengths[2])

			if length1+length2 > length3 && length1+length3 > length2 && length2+length3 > length1 {
				count++
			}
		}
		fmt.Println(count)
	}
}
