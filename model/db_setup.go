package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DB_Connection() *sql.DB {

	err := godotenv.Load("C:/Users/smc/go/src/github.com/users/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbUrl := fmt.Sprintf("%s:%s@tcp(localhost:3306)/users", dbUser, dbPassword)

	db, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		log.Println("Error connecting to database", err.Error())
		return nil
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging database", err.Error())
		return nil
	}
	return db
}
