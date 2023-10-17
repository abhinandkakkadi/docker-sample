package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func main() {
	
  tpl = template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}
	
	tpl.ExecuteTemplate(w, "index.html", nil)
}