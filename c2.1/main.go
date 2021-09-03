package main

import (
	"html/template"
	"net/http"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("c2.1/index.tmpl")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}
