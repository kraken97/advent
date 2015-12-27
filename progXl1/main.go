package main

import (
	"fmt"
	"github.com/tealeg/xlsx"

	"io"
	"net/http"
	"strconv"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<html>
<head>
    <meta charset="utf-8">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
    <title>XLSX</title>
    <style type="text/css">
        input,button {
            margin-top: 10px;
        }      
    </style>
</head>
<body>
    <form action="/save" method="post">
    <div class="col-sm-8">
    <input type="text" class="form-control" placeholder="Файл" name="file1"></input>
    </div>
    <div class="col-sm-4">
    <input type="text" class="form-control" placeholder="Колонка" name="col1"></input>
    </div>
    <div class="col-sm-8">
    <input type="text" class="form-control" placeholder="Файл" name="file2"></input>
    </div>
    <div class="col-sm-4">
    <input type="text" class="form-control" placeholder="Колонка" name="col2"></input>
    </div>
    <center><input type="submit" class="btn btn-success" value="Сравнить"></input></center>
    </form>
</body>
</html>`,
	)
	file1 := r.FormValue("file1")
	file2 := r.FormValue("file2")
	col1 := r.FormValue("col1")
	col2 := r.FormValue("col2")

}
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9000", nil)
}

func calculate(file1 string, file2 string, col1 string, col2 string) {

	// link
	excelFileName := file1
	excelFileName2 := file2
	//open
	xlFile1, error1 := xlsx.OpenFile(excelFileName)
	xlFile2, error2 := xlsx.OpenFile(excelFileName2)
	//error
	if error1 != nil || error2 != nil {
		fmt.Println("file open errr")
	}
	sh1, sh2 := xlFile1.Sheets[0], xlFile2.Sheets[0]
	findEqual(sh1, sh2, col1, col2)

}
func findEqual(sh1 *xlsx.Sheet, sh2 *xlsx.Sheet, col1 string, col2 string) {
	col1 = strconv.Atoi(col1)
	col2, _ = strconv.Atoi(col2)
	searchCol := col1

	for i := 0; i < sh1.MaxRow; i++ {
		sh1YACHEYKA := sh1.Cell(i, searchCol)

		for k := 0; k < sh2.MaxRow; k++ {
			fn1Yacheyka := sh2.Cell(k, col2)
			if sh1YACHEYKA.Value == fn1Yacheyka.Value && fn1Yacheyka.Value != "" {
				fmt.Println(i, "--", sh1YACHEYKA, "--", k, fn1Yacheyka)
			}
		}
	}
}
