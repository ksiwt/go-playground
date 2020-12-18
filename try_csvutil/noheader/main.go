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
	ID   int
	Name string
	Age  int `csv:",omitempty"`
	City string
}

func main() {

	csvReader := csv.NewReader(strings.NewReader(`
1,John,27,la
2,Bob,,ny`))

	// in real application this should be done once in init function.
	userHeader, err := csvutil.Header(User{}, "csv")
	if err != nil {
		log.Fatal(err)
	}

	dec, err := csvutil.NewDecoder(csvReader, userHeader...)
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	for {
		var u User
		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	fmt.Printf("%+v", users)
}

// output:
// [{ID:1 Name:John Age:27 City:la} {ID:2 Name:Bob Age:0 City:ny}]
