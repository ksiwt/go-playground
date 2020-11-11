package main

import (
	"os"
	"text/template"
)

func main() {
	text := "My Name is {{.name}} and {{.age}} years old"
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		panic(err)
	}

	object := map[string]interface{}{
		"name": "Taro",
		"age":  30,
	}

	err = tmpl.Execute(os.Stdout, object) // My Name is Taro and 30 years old
	if err != nil {
		panic(err)
	}
}
