package main

import (
	"database/sql"
	"fmt"
	"log"
    "example.com/m/pkg/user"

	_ "github.com/lib/pq"
)

func main() {
	// Set up the database connection.
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the user service.
	service := user.NewService(db)

	// Create a new user.
	newUser := user.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	if err := service.CreateUser(newUser); err != nil {
		log.Fatal(err)
	}

	// Get all users.
	users, err := service.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	// Print all users.
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}
}
