package main

import "html/template"

func main() {
	template.ParseFiles("index.tmpl", "index2.tmpl")

}
