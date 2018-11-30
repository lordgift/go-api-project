package service

import (
	"github.com/gin-gonic/gin"
	"bank-account/persistence"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct {
	db                 *sql.DB
	userService        persistence.UserService
	bankAccountService persistence.BankAccountService
}

func InitiateDB() *Server {
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
		userService: &persistence.UserServiceImp{
			DB: db,
		},
		bankAccountService: &persistence.BankAccountServiceImp{
			DB: db,
		},
	}

	return s
}

func SetupRoute(s *Server) *gin.Engine {
	r := gin.Default()

	root := r.Group("/")
	root.GET("/users", s.All)
	root.GET("/users/:id", s.All)
	root.POST("/users", s.All)
	root.PUT("/users/:id", s.All)
	root.DELETE("/users/:id", s.All)
	root.POST("/users/:id/bankAccounts", s.All)
	root.GET("/users/:id/bankAccounts", s.All)
	root.DELETE("/bankAccounts/:id", s.All)
	root.PUT("/bankAccounts/:id/withdraw", s.All)
	root.PUT("/bankAccounts/:id/deposit", s.All)
	root.POST("/transfers", s.All)

	// - GET /users สำหรับ list users ทั้งหมด
	// - GET /users/:id สำหรับ ดึงข้อมูล user ตาม id
	// - POST /users สำหรับสร้าง user ใหม่
	// - request body: {“ﬁrst_name”: “John”, “last_name”: “Doe”}
	// - PUT /users/:id สำหรับ update user ตาม id
	// - request body: {“ﬁrst_name”: “John”, “last_name”: “Doe”}
	// - DELETE /users/:id สำหรับลบ user ตาม id
	// - POST /users/:id/bankAccounts สำหรับเพิ่ม bank account ให้ user ตาม :id

	// - request body: {“account_number”: “123456”, “name”: “John Doe”}
	// - GET /users/:id/bankAccounts สำหรับดึง bank account ทั้งหมดของ user ตาม :id
	// - DELETE /bankAccounts/:id สำหรับลบ bank account ตาม :id
	// - PUT /bankAccounts/:id/withdraw สำหรับถอนเงินจาก bank account ตาม :id
	// - request body: {“amount”: 100000}
	// - PUT /bankAccounts/:id/deposit สำหรับฝากเงินเข้า bank account ตาม :id
	// - request body: {“amount”: 100000}
	// - POST /transfers เป็นการโอนเงิน โดยรับ request body JSON ดังนี้
	// {amount: 100000,  “from”: [accoun_number], “to”: [account_number]}

	return r
}

func (s *Server) All(c *gin.Context) {
	todos, err := s.userService.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}
