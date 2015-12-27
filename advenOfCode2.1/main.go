// advenOfCode2.1 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	path := "C:\\gowork\\src\\advenOfCode2.1\\text.txt"
	File, _ := readLines(path)
	var sum int64 = 0
	var res int64 = 0
	for line := range File {
		a, b, c := parse(File[line])
		res = calculateSum(a, b, c)
		sum += res
		fmt.Println(line, sum)
	}
	fmt.Println(sum)

}

func parse(str string) (a int64, b int64, c int64) {
	reg, _ := regexp.Compile("[x]")
	strM := reg.Split(str, -1)
	a, _ = strconv.ParseInt(strM[0], 0, 64)
	b, _ = strconv.ParseInt(strM[1], 0, 64)
	c, _ = strconv.ParseInt(strM[2], 0, 64)

	return
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func calculateSum(l int64, w int64, h int64) (b int64) {
	a, b := smallest(l, w, h)
	return 2*w*h + 2*h*l + 2*l*w + a*b
}

func smallest(l int64, w int64, h int64) (t int64, y int64) {
	if l <= h && w <= h {
		return l, w
	}
	if h <= w && l <= w {
		return l, h
	} else {
		return w, h
	}

}
