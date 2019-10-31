package main

import (
	"bank-account/service"
	"os"

	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	os.Setenv("DATABASE_URL", "root:password@tcp(localhost:3306)/testdb")
	os.Setenv("PORT", "8000")

	s := service.InitiateDB()
	r := service.SetupRoute(s)

	r.Run(":" + os.Getenv("PORT"))
}
