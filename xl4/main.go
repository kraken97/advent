// xl4 project main.go
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	// link
	excelFileName := "C:\\gowork\\src\\xl4\\test1.xlsx"
	excelFileName2 := "C:\\gowork\\src\\xl4\\test2.xlsx"
	//open
	xlFile1, error1 := xlsx.OpenFile(excelFileName)
	xlFile2, error2 := xlsx.OpenFile(excelFileName2)
	//error
	if error1 != nil || error2 != nil {
		fmt.Println("file open errr")
	}
	sh1, sh2 := xlFile1.Sheets[0], xlFile2.Sheets[0]
	findEqual(sh1, sh2)

}

func findEqual(sh1 *xlsx.Sheet, sh2 *xlsx.Sheet) {
	searchCol := 4

	for i := 0; i < sh1.MaxRow; i++ {
		sh1YACHEYKA := sh1.Cell(i, searchCol)

		for k := 0; k < sh2.MaxRow; k++ {
			fn1Yacheyka := sh2.Cell(k, 3)
			if sh1YACHEYKA.Value == fn1Yacheyka.Value && fn1Yacheyka.Value != "" {
				fmt.Println(i, "--", sh1YACHEYKA, "--", k, fn1Yacheyka)
			}
		}
	}
}
