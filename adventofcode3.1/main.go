// adventofcode3.1 project main.go
package main

import (
	"fmt"

	"io/ioutil"
)

const X = 1000
const Y = 1000

func main() {

	path := "C:\\gowork\\src\\adventofcode3.1\\text.txt"
	File := readLines(path)
	var x1 int = X / 2
	var y1 int = Y / 2
	var x2 int = X / 2
	var y2 int = Y / 2
	var pole [Y][X]int

	for i := 0; i < len(File); i++ {
		if File[i] == '>' {
			x1 += 1
		} else if File[i] == '^' {

			y1 += 1
		} else if File[i] == 'v' {
			y1 -= 1
		} else if File[i] == '<' {
			x1 -= 1
		}
		pole[y1][x1] += 1
		i++
		if File[i] == '>' {
			x2 += 1
		} else if File[i] == '^' {

			y2 += 1
		} else if File[i] == 'v' {
			y2 -= 1
		} else if File[i] == '<' {
			x2 -= 1
		}
		pole[y2][x2] += 1
	}
	fmt.Println(check(pole))

}

func readLines(path string) string {
	//	file, err := os.Open(path)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer file.Close()
	//	stat, err := file.Stat()
	//	if err != nil {
	//		panic(err)
	//	}

	//	bs := make([]byte, stat.Size())
	//	_, err = file.Read(bs)
	//	if err != nil {
	//		panic(err)
	//	}

	//	str := string(bs)
	////	fmt.Println(str)

	dat, _ := ioutil.ReadFile(path)
	fmt.Print(string(dat))
	return string(dat)
}

func check(pole [X][Y]int) int {
	count := 0
	for i := 0; i < Y; i++ {
		for k := 0; k < X; k++ {
			if pole[i][k] >= 1 {
				count++
			}
		}
	}
	return count
}
