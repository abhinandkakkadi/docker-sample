package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func main() {
	
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}
	
	w.Write([]byte("OK"))
}