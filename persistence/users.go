package persistence

import "database/sql"

type User struct {
	ID        int64  `json:"id"`
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
}

type UserService interface {
	All() ([]User, error)
	// 	Insert(todo *Todo) error
	// 	GetByID(id int) (*Todo, error)
	// 	DeleteByID(id int) error
	// 	Update(id int, body string) (*Todo, error)
}

type UserServiceImp struct {
	DB *sql.DB
}

func (s *UserServiceImp) All() ([]User, error) {
	rows, err := s.DB.Query("SELECT id, first_name, last_name FROM Users")
	if err != nil {
		return nil, err
	}
	users := []User{} // set empty slice without nil
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.firstName, &user.lastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
