// ор project ор.go
package ор
import (
   "github.com/tealeg/xlsx"
    "fmt"
   
)
func main(){
	  fmt.Print("kk")
	
	excelFileName := "D:\test.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
	    if err != nil {
        fmt.Errorf("file not found")
    }
    sheet :=  xlFile.Sheets[0] 
         row  sheet.Rows[2] 
            for _, cell := range row.Cells {
                fmt.Printf("%s\n", cell.String())
            }
        
    
	
}