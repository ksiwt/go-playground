package main

import (
	"fmt"
	"time"

	"github.com/jszwec/csvutil"
)

type Address struct {
	City string
	Country string
}

type User struct {
	Name  string
	Address
	Age   int `csv:"age,omitempty"`
	CreatedAt time.Time
}

func main() {
	users := []User{
		{
			Name:      "John",
			Address:   Address{"Boston", "USA"},
			Age:       26,
			CreatedAt: time.Date(2010, 6, 2, 12, 0, 0, 0, time.UTC),
		},
		{
			Name:    "Alice",
			Address: Address{"SF", "USA"},
		},
	}

	u, err := csvutil.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(u))
}

// output:
//Name,City,Country,age,CreatedAt
//John,Boston,USA,26,2010-06-02T12:00:00Z
//Alice,SF,USA,,0001-01-01T00:00:00Z