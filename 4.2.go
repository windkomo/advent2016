package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sort"
)

func DecryptRoomName (s string, sectorId int) string {
	//a 97 z 122
	s = strings.Replace(s, "-", " ", -1)
	result := ""
	rotation := sectorId % 26

	for _, char := range s {
		if char != 32 {
			if int(char) + rotation > 122 {
				result += string(97 - 1 + int(char) + rotation - 122)
			} else {
				result += string(int(char) + rotation)
			}
		} else {
			result += " "
		}
  }

	return result
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func makeChecksum(s string) string {
	occs := make(map[string]int)
	results := ""

	for i := 0; i < len(s); i++ {
		charCount := strings.Count(s, string(s[i]))
		occs[string(s[i])] = charCount
	}

	// Sorted keys : letters that have the high count first
	var keys []string
		for k := range occs {
				keys = append(keys, k)
		}

		sort.Strings(keys)

	// Sorted valued : sorted like keys
	var values []int
    for _,v := range keys {
        values = append(values, occs[v])
    }

    sort.Sort(sort.Reverse(sort.IntSlice(values)))

	for i := 0; i < 5; i++ {
		highCount := values[i]
		for _, v := range keys {
			if highCount == occs[v] && SliceIndex(len(results), func(i int) bool { return string(results[i]) == v }) == -1 {
				results += v
				break
			}
    }
	}

	return results
}

func main() {
	DecryptRoomName("qzmt-zixmtkozy-ivhz", 343)

	re := regexp.MustCompile("(([a-z]+-)+)(\\d+)\\[([a-z]+)\\]")
	if file, err := os.Open("input.4.txt"); err == nil {

		// make sure it gets closed
		defer file.Close()

		scanner := bufio.NewScanner(file)
		sum := 0

		for scanner.Scan() {

			parts := re.FindAllStringSubmatch(scanner.Text(), -1)[0]
			// fmt.Println(parts)

			name := parts[1]
			cleanName := SortString(strings.Replace(name, "-", "", -1))
			sectorId, _ := strconv.Atoi(parts[3])
			decryptedName := DecryptRoomName(name, sectorId)
			fmt.Println(decryptedName, sectorId)
			checksum := parts[4]
			testChecksum := makeChecksum(cleanName)

			if testChecksum == checksum {
				sum += sectorId
			}
		}
	}
}
