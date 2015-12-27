// adventofcode project main.go
package main

import (
	"fmt"
)

func main() {
	incby1k()
}

var (
	j int = 0
	z int = 1
)

func findNumbers(n int) int {
	sum := 0
	if j == 50 {
		j = 0
		z++
	}
	for i := z; i <= n; i++ {
		if n%i == 0 {

			sum += i
		}
	}
	j++
	return sum * 11
}
func incby1k() {
	var num int = 0
	i := 616020
	for ; num <= 34000000; i += 1 {

		num = findNumbers(i)

		fmt.Println(i, num)

	}

}
