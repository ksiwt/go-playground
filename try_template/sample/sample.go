package main

import (
	"os"
	"text/template"
)

func main() {
	text := "Hello {{.}}"
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, "Taro") // Hello Taro
	if err != nil {
		panic(err)
	}
}
