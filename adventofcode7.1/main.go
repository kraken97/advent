// adventofcode7.1 project main.go
package main

import (
	"fmt"
)

func main() {
	str := "11132221"
	for i := 0; i < 40; i++ {
		str = check(str)
		fmt.Println(str)

	}

}

func check(str string) string {
	count := 1

	for i := 0; i < len(str)-1; i++ {
		for k := i + 1; k < len(str); k++ {
			if str[i] != str[k] {
				i = k - 2

				break
			}
			if str[i] == str[k] {
				count++

			}

		}

		fmt.Println(count)
		count = 1
	}
	return ""
}
