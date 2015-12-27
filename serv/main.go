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
    <form action="save" method="post">
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

}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/save", result)
	http.ListenAndServe(":9000", nil)
}
func result(res http.ResponseWriter, req *http.Request) {
	file1 := req.FormValue("file1")
	file2 := req.FormValue("file2")
	col1 := req.FormValue("col1")
	col2 := req.FormValue("col2")
	//		reg := regexp.MustCompile("\\")
	//		col1 = reg.ReplaceAllString(col1, "\\\\")
	////		col2 = reg.ReplaceAllString(col2, "\\\\")
	io.WriteString(res, calculate(file1, file2, col1, col2))
}

func calculate(file1 string, file2 string, col1 string, col2 string) string {

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
	return findEqual(sh1, sh2, col1, col2)

}
func findEqual(sh1 *xlsx.Sheet, sh2 *xlsx.Sheet, col1 string, col2 string) string {
	c1, _ := strconv.Atoi(col1)
	c2, _ := strconv.Atoi(col2)
	searchCol := c1
	res := ""
	for i := 0; i < sh1.MaxRow; i++ {
		sh1YACHEYKA := sh1.Cell(i, searchCol)

		for k := 0; k < sh2.MaxRow; k++ {
			fn1Yacheyka := sh2.Cell(k, c2)
			if sh1YACHEYKA.Value == fn1Yacheyka.Value && fn1Yacheyka.Value != "" {
				res += fmt.Sprint(i, "--", sh1YACHEYKA, "--", k, fn1Yacheyka, "\n")
			}
		}
	}
	fmt.Println(res)
	return res
}
