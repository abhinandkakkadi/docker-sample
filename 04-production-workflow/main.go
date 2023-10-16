package main

import (
	"net/http"
	"text/template"
)

var tpl template.Template

func init() {
	
}

func main() {
	var tpl = template.Must(template.New("sample").ParseFiles("templates/index.html"))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}

	tpl.ExecuteTemplate(w, "index.html", nil)
}