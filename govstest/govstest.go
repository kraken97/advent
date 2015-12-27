package main

import (
	"fmt"
	"sort"
)


func main() {
     x:=[]int{2,4,6,65,6,7,8,6,5,5,7,6}

	fmt.Println(min(x))
}

func min(a []int) int {
	min:=-10000
	for i:=0;i<a.lent();i++{
		if [i]a<min {
			min= [i]a
		}
	}
}
