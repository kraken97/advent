// xl; project main.go
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	fmt.Print("kk")

	excelFileName := "D:\test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Print("kk")
	}
	for err1, sheet := range xlFile.Sheets {
		if err1 != nil {
			fmt.Print("kk1")
		}
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Printf("%s\n", cell.String())
			}
		}
	}

}
