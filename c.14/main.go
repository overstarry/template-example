package main

import "html/template"

const goTemplate = `
package main 

import "fmt"

func main() {
	fmt.Println("hello")
}
`

func main() {
	t, err := template.New("go").Parse(goTemplate)
}
