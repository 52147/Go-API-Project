package user

import (
	"database/sql"
)

// Service provides methods to interact with the user data.
type Service struct {
	db *sql.DB
}

// NewService creates a new instance of the user service.
func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

// GetAll retrieves all users from the database.
func (s *Service) GetAll() ([]User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// GetSingle retrieves a single user by ID from the database.
func (s *Service) GetSingle(userID int) (*User, error) {
	var u User
	row := s.db.QueryRow("SELECT * FROM users WHERE id = $1", userID)
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &u, nil
}

// Create adds a new user to the database.
func (s *Service) Create(newUser User) (*User, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", newUser.Name, newUser.Email).Scan(&id)
	if err != nil {
		return nil, err
	}

	newUser.ID = id
	return &newUser, nil
}

// Update updates an existing user in the database.
func (s *Service) Update(userID int, updatedUser User) (*User, error) {
	_, err := s.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", updatedUser.Name, updatedUser.Email, userID)
	if err != nil {
		return nil, err
	}

	updatedUser.ID = userID
	return &updatedUser, nil
}

// Delete removes a user from the database.
func (s *Service) Delete(userID int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", userID)
	return err
}
