package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "reader"
	DB_PASSWORD = "pass"
	DB_NAME     = "postgres"
)

func GetConnection(dbname string) sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return *db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func liner(token string, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(token)
	}
	fmt.Println()
}
