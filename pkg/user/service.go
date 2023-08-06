package user

import "database/sql"

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

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

func (s *Service) CreateUser(user User) error {
	_, err := s.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}
