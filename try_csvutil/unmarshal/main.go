package main

import (
	"fmt"
	"time"

	"github.com/jszwec/csvutil"
)

type User struct {
	ID        uint64 `csv:"id"`
	Name      string `csv:"name"`
	Age       int
	CreatedAt time.Time
}

func main() {
	var csvInput = []byte(`
id,name,Age,CreatedAt
1,jacek,20,2012-04-01T15:00:00Z
2,john,30,0001-01-01T00:00:00Z`,
	)

	var users []User
	if err := csvutil.Unmarshal(csvInput, &users); err != nil {
		fmt.Println("error:", err)
	}

	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

// output:
// {ID:1 Name:jacek Age:20 CreatedAt:2012-04-01 15:00:00 +0000 UTC}
// {ID:2 Name:john Age:30 CreatedAt:0001-01-01 00:00:00 +0000 UTC}
