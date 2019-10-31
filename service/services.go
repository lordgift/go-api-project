package service

import (
	"bank-account/persistence"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	db                 *sql.DB
	userService        persistence.UserService
	bankAccountService persistence.BankAccountService
}

func InitiateDB() *Server {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	// createTable := `
	// CREATE TABLE IF NOT EXISTS Users (
	// 	id int unsigned auto_increment primary key,
	// 	first_name VARCHAR(100),
	// 	last_name VARCHAR(100)
	// );
	// CREATE TABLE IF NOT EXISTS BankAccounts (
	// 	id int unsigned auto_increment primary key,
	// 	user_id VARCHAR(100),
	// 	account_number VARCHAR(100),
	// 	name VARCHAR(100),
	// 	balance FLOAT
	// );
	// `
	// if _, err := db.Exec(createTable); err != nil {
	// 	log.Fatal(err)
	// }

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
	root.GET("/users/:id", s.FindByID)
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
	users, err := s.userService.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (s *Server) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := s.userService.FindByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
