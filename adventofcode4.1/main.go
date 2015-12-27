// adventofcode4.1 project main.go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	for i := 1; i < 99999; i++ {
		str := "yzbqklnj" + strconv.Itoa(i)

		go GetMD5Hash(str, i)

	}
	end := time.Now()
	fmt.Println(now, end)

}

func GetMD5Hash(text string, i int) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	if i%1000 == 0 {
		fmt.Println(i)

	}
	return hex.EncodeToString(hasher.Sum(nil))

}
