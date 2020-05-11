package dba

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "reader"
	dbPassword = "pass"
	dbName     = "postgres"
)

func GetConnection() sqlx.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		getEnv("DB_HOST", dbHost),
		getEnv("DB_PORT", dbPort),
		getEnv("DB_NAME", dbName),
		getEnv("DB_USER", dbUser),
		getEnv("DB_PASSWORD", dbPassword),
	)
	db, err := sqlx.Open("postgres", dbinfo)
	CheckErr(err)
	return *db
}

func CheckErr(err error) {
	if err != nil {
		liner("=", 64)
		panic(err)
		liner("-", 64)
	}
}

func liner(token string, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(token)
	}
	fmt.Println()
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
