package main

import (
	"html/template"
	"log"
)

func main() {
	_, err := template.New("index").Parse("index.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

}
