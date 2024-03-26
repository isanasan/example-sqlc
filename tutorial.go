package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"example-sqlc/db"
)

func run() error {
	ctx := context.Background()

	conn, err := sql.Open("mysql", "root@/test?parseTime=true")
	if err != nil {
		return err
	}

	queries := db.New(conn)

	// create an user
	result, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:  "test",
		Email: "test@test.com",
		Age:   20})
	if err != nil {
		return err
	}

	insertedUserID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedUserID)

	// get the author we just inserted
	fetchedUser, err := queries.GetUser(ctx, uint64(insertedUserID))
	if err != nil {
		return err
	}

	log.Println(fetchedUser)

	// update an user age
	queries.UpdateUserAges(ctx, db.UpdateUserAgesParams{
		Age: 40,
		ID:  uint64(insertedUserID)})

	// get the author we just inserted
	fetchedUpdatedUser, err := queries.GetUser(ctx, uint64(insertedUserID))
	if err != nil {
		return err
	}

	log.Println(fetchedUpdatedUser)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
