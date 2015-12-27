// adventofcode20 project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	печать := fmt.Print
	открыть := os.Open

	результат, уй := открыть("C:\\test.txt")
	if уй != nil {
		panic(уй)
	} else {
		печать(результат)
	}
}
