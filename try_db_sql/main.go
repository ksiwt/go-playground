package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"go_playground/try_db_sql/repository"
)

func main() {
	db, err := sql.Open("mysql",
		"root:mysql@tcp(127.0.0.1:43306)/test",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(db)

	ctx := context.Background()
	animals, err := repo.FindAllAnimals(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, a := range animals {
		fmt.Printf("ID: %v, Name: %v \n", a.ID, a.Name)
	}
}
