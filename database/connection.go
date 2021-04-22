package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func dsn(db string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUsername, dbPassword, dbHost, dbName)
}

func Dberror(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(os.Getenv("DB_NAME")))
	Dberror(err)
	return db, err
}
