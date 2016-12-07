package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	doorId := "ugkcyxxp"
	password := make([]string, 8)
	valuesCount := 0
	i := 0

	for valuesCount < 8 {
		s := doorId + strconv.Itoa(i)
		hash := md5.Sum([]byte(s))
		hashString := hex.EncodeToString(hash[:])

		if hashString[:5] == "00000" {
			position, err := strconv.Atoi(string(hashString[5]))

			if err == nil && position < 8 && len(password[position]) == 0 {
				password[position] = string(hashString[6])
				valuesCount++
			}
		}

		i++
	}

	fmt.Println(strings.Join(password, ""))
}
