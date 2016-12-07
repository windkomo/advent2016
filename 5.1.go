package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	doorId := "ugkcyxxp"
	password := ""
	i := 0

	for len(password) < 8 {
		s := doorId + strconv.Itoa(i)
		hash := md5.Sum([]byte(s))
		hashString := hex.EncodeToString(hash[:])

		if hashString[:5] == "00000" {
			password += string(hashString[5])
		}

		i++
	}

	fmt.Println(password)
}
