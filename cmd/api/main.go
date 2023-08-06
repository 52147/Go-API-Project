package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"example.com/m/pkg/user"
    "github.com/gorilla/mux"

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
		Email: "john.doe12345@example.com",
	}
	createdUser, err := service.Create(newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New User created: ID: %d, Name: %s, Email: %s\n", createdUser.ID, createdUser.Name, createdUser.Email)

	// Get all users.
	users, err := service.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	// Print all users.
	fmt.Println("All Users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}

	// Create the handler and set up routes.
    h := user.NewHandler(db)

    r := mux.NewRouter()
    r.HandleFunc("/users", h.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", h.GetSingleUser).Methods("GET") // Handle the "/users/{id}" route.
    r.HandleFunc("/users", h.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")

    fmt.Println("Starting server on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
