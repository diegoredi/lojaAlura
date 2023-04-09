package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	user := "postgres"
	password := "135790"
	dbname := "alura_loja"
	host := "localhost"
	port := "5432"

	sc := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

	db, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err.Error())
	}
	return db
}
