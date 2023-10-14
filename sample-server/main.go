package main

import "net/http"



func main() {

	http.HandleFunc("/",index)
	http.ListenAndServe(":3000",nil)
	
}

func index(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte("OK"))

}