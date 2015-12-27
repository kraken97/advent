// WebSite project main.go
package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./files"))))
	http.ListenAndServe(":8000", nil)
}
