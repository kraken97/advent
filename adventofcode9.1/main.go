// adventofcode9.1 project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	// create array
	res := readFile("D:\\test.txt")
	re := regexp.MustCompile("[^0-9]")
	res = re.ReplaceAllString(res, " ")
	s := regexp.MustCompile(" {1,}").Split(res, -1)
	var matrix [8][8]int
	z := 1
	//matrix
	for i := 0; i < 8; i++ {
		for k := i; k < 8; k++ {
			if i == k {
				matrix[i][k] = 0
			} else {
				matrix[i][k], _ = strconv.Atoi(s[z])
				matrix[k][i], _ = strconv.Atoi(s[z])
				z++
			}

		}

	}
	//min := 0
	for k := range matrix {
		fmt.Println(matrix[k])
	}
	getMin(matrix)

}
func readFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}
func getMin(matrix [8][8]int) {
}
