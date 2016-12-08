
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if file, err := os.Open("input.6.txt"); err == nil {
		defer file.Close()

    store := make([]map[rune]int, 8)
    results := make([]string, 8)
    counts := make([]int, 8)

    for i := 0; i < 8 ; i++ {
      store[i] = make(map[rune]int, 26)
    }

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
      line := scanner.Text()
      for index, char := range line {
          store[index][char]++
      }
		}

    for i := 0; i < 8 ; i++ {
      for char, count := range store[i] {
        if counts[i] < count {
          results[i] = string(char)
          counts[i] = count
        }
      }
    }

    fmt.Println(strings.Join(results, ""))
	}
}
