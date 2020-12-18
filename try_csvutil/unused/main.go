package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/jszwec/csvutil"
)

type User struct {
	Name      string            `csv:"name"`
	City      string            `csv:"city"`
	Age       int               `csv:"age"`
	OtherData map[string]string `csv:"-"`
}

func main() {
	csvReader := csv.NewReader(strings.NewReader(`
name,age,city,zip
alice,25,la,90005
bob,30,ny,10005`))

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	header := dec.Header()
	var users []User
	for {
		u := User{OtherData: make(map[string]string)}

		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for _, i := range dec.Unused() {
			u.OtherData[header[i]] = dec.Record()[i]
		}
		users = append(users, u)
	}

	fmt.Println(users)
}

// output:
// [{alice la 25 map[zip:90005]} {bob ny 30 map[zip:10005]}]