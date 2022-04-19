package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "costanatilic"
)

func Dbconfig() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Successfully connected!")
}
