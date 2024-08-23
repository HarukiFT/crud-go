package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const connectionStringTemplate string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"

func GetDB() *sqlx.DB {
	connString := fmt.Sprintf(connectionStringTemplate, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	return db
}
