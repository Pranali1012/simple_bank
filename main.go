package main

import (
	"database/sql"
	"log"

	db "github.com/pranali1012/simple_bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	//Connecting to the database
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
	}
	defer conn.Close()

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server := api.newServer

	//Creating an instance of a Queries struct from sqlc package
	//queries := sqlc.New(conn)

	// //To create an account
	// //Creating an instance of a CreateAccountParams struct from sqlc package
	// accountParams := sqlc.CreateAccountParams{Owner: "Pranali", Balance: 35000, Currency: "INR"}

	// account, err := queries.CreateAccount(context.Background(), accountParams)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Created account with ID: %d\n", account.ID)

	// //To create an entry
	// //Creating an instance of a CreateEntryParams struct from sqlc package
	// entryParams := sqlc.CreateEntryParams{AccountID: account.ID, Amount: 12343}

	// entry, err := queries.CreateEntry(context.Background(),entryParams)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Created an entry with account ID: %d\n", entry.AccountID)

	// // //To update an entry
	// // //Creating an instance of a UpdateEntryParams struct from sqlc package
	// updateEntryParams := sqlc.UpdateEntryParams{ID: 1, Amount: 6753}

	// updatedEntry,err := queries.UpdateEntry(context.Background(),updateEntryParams)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("updated an entry with account ID: %d\n", updatedEntry.ID)

	// // Created an amount transfer between two accounts and updated the account balance
	// createTransferParams := sqlc.CreateTransferParams{FromAccountID: 4, ToAccountID: 2, Amount: 2000}

	// transfer, err := queries.CreateTransfer(context.Background(), createTransferParams)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("created transfer  with account ID: %d\n", transfer.ID)

	// fromAccount, err := queries.GetAccount(context.Background(), transfer.FromAccountID)
	// if err != nil {
	// 	panic(err)
	// }

	// //To update an account
	// //Creating an instance of a UpdateAccountParams struct from sqlc package
	// updateAccountParams := sqlc.UpdateAccountParams{ID: transfer.FromAccountID, Balance: fromAccount.Balance - transfer.Amount}

	// updatedAccount, err := queries.UpdateAccount(context.Background(), updateAccountParams)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("updated an entry with account ID: %d\n", updatedAccount.Balance)

	//To Delete an account
	//Creating an instance of a DeleteAccountParams strunct from sqlc package

	// err = queries.DeleteAccount(context.Background(), 5)
	// if err != nil {
	// 	panic(err)
	// }

	// //To update an entry
	// //Creating an instance of a UpdateEntryParams struct from sqlc package
	// updateEntryParams := sqlc.UpdateEntryParams{ID: 5, Amount: 0}

	// updatedEntry, err := queries.UpdateEntry(context.Background(), updateEntryParams)
	// if err != nil {
	// 	panic(err)
	// }
	// if updatedEntry.Amount == 0 {
	// 	err = queries.DeleteAccount(context.Background(), 5)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}
