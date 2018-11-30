package persistence

import (
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserService interface {
	All() ([]User, error)
	FindByID(id int) (*User, error)
	// 	Insert(todo *Todo) error
	// 	GetByID(id int) (*Todo, error)
	// 	DeleteByID(id int) error
	// 	Update(id int, body string) (*Todo, error)
}

type UserServiceImp struct {
	DB *sql.DB
}

func (s *UserServiceImp) All() ([]User, error) {
	rows, err := s.DB.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	users := []User{} // set empty slice without nil
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *UserServiceImp) FindByID(id int) (*User, error) {
	row := s.DB.QueryRow("SELECT * FROM Users WHERE id = $1", id)
	var user User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
