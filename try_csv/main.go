package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	read(in)
	readAll(in)

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	write(records)
	writeAll(records)
}

// 1行ずつ読み込む。
func read(in string) {
	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(record)
	}
}

// 一度に全て読み込む。
func readAll(in string) {
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}

func write(records [][]string) {
	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}

		// バッファリングされた出力をフラッシュする。
		w.Flush()

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
	}
}

func writeAll(records [][]string) {
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // フラッシュは内部で呼び出される。

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
