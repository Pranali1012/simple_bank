package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pranali1012/simple_bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}

	defer conn.Close()

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(conn)

	accountParams := sqlc.CreateAccountParams{Owner: "Pratik", Balance: 25000, Currency: "INR"}

	account, err := queries.CreateAccount(context.Background(), accountParams)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created account with ID: %d\n", account.ID)
}
