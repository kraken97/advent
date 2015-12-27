package main

import "fmt"
import "main_test/pack"

//какр какр
func main() {
	xs := []float64{1, 2, 3, 4}
	avg := pack.Average(xs)
	fmt.Println(avg)
}
