package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func handlerPost(responseWriter http.ResponseWriter, response *http.Request) {

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"), os.Getenv("sslmode"))

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	id := 1
	resultSet, err := db.Query("SELECT acount_name FROM accounts WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer resultSet.Close()

	var name string
	for resultSet.Next() {
		if err := resultSet.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name of the id:%d is %s\n", id, name)
	}

	type userPosted struct {
		Mensagem string `json:"msg"`
	}

	respondWithJson(responseWriter, 200, userPosted{Mensagem: name})
}
