package dba

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "reader"
	DB_PASSWORD = "pass"
	DB_NAME     = "postgres"
)

func GetConnection(dbname string) sqlx.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, dbname)
	db, err := sqlx.Open("postgres", dbinfo)
	CheckErr(err)
	return *db
}

func CheckErr(err error) {
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
