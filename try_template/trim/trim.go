package main

import (
	"os"
	"text/template"
)

func main() {
	text := "吾輩は猫である。名前はまだ無い。\n {{- . }}" +
		"{{ . -}} どこで生れたかとんと見当がつかぬ。\n" +
		"何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。"
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, "")
	// 吾輩は猫である。名前はまだ無い。どこで生れたかとんと見当がつかぬ。
	// 何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。
	if err != nil {
		panic(err)
	}
}
