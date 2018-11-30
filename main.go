package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	db *sql.DB
}

func setupRoute(s *Server) *gin.Engine {
	r := gin.Default()
	return r
}

func main() {

	os.Setenv("DATABASE_URL", "postgres://postgres:gotraining@localhost/postgres?sslmode=disable")
	os.Setenv("PORT", "8000")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS Users (
		id SERIAL PRIMARY KEY,
		first_name TEXT,
		last_name TEXT
	);
	CREATE TABLE IF NOT EXISTS BankAccounts (
		id SERIAL PRIMARY KEY,
		user_id TEXT,
		account_number TEXT,
		name TEXT,
		balance FLOAT
	);
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	s := &Server{
		db: db,
	}

	r := setupRoute(s)

	r.Run(":" + os.Getenv("PORT"))
}
