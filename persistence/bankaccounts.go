package persistence

import "database/sql"

type BankAccount struct {
	ID            int64   `json:"id"`
	userID        string  `json:"user_id"`
	accountNumber string  `json:"account_number"`
	name          string  `json:"name"`
	balance       float64 `json:"balance"`
}

type BankAccountService interface {
	// Insert(s *Secret) error
}

type BankAccountServiceImp struct {
	DB *sql.DB
}
