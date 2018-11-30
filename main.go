package main

import (
	"bank-account/service"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	os.Setenv("DATABASE_URL", "postgres://postgres:gotraining@localhost/postgres?sslmode=disable")
	os.Setenv("PORT", "8000")

	s := service.InitiateDB()
	r := service.SetupRoute(s)

	r.Run(":" + os.Getenv("PORT"))
}
