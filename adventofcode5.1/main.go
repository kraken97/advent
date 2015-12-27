// adventofcode5.1 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {
	file := openFile("C:\\gowork\\src\\adventofcode5.1\\text.txt")
	//	vowels := [5]string{
	//		"a",
	//		"e",
	//		"i",
	//		"o",
	//		"u",
	//	}
	//	notContain := [4]string{
	//		"ab",
	//		"cd",
	//		"pq",
	//		"xy",
	//	}
	count := 0
	//file := [1]string{"uurcxstgmygtbstg"}

	for i := range file {

		b := findthree(file[i])
		c := findtwo(file[i])

		if b && c {
			count++
		}
		fmt.Println(b, c, file[i], (b && c))

	}

	fmt.Println(count)

}

//func cont1(vowels [5]string, str string) bool {
//	count := 0
//	for i := range vowels {

//		count += strings.Count(str, vowels[i])

//	}
//	if count >= 3 {
//		return true
//	} else {
//		return false
//	}
//}
//func cont(vowels [4]string, str string) bool {

//	for i := range vowels {
//		if strings.Contains(str, vowels[i]) {
//			return false
//		}
//	}
//	return true
//}

func openFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

//5.2 если хх и хх
func findtwo(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		strf := str[i : i+2]

		for k := i + 2; k < len(str)-1; k++ {
			if strf == str[k:k+2] {
				return true

			}
		}

	}

	return false
}

//5.2 если хух
func findthree(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}
	return false
}
