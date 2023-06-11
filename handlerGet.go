package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handlerGet(responseWriter http.ResponseWriter, request *http.Request) {

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"), os.Getenv("sslmode"))

	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(request.URL.RawQuery)
	id := 1
	rows, err := db.Query("SELECT acount_name FROM accounts WHERE id = $1", id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var name string
	var quant int16
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name of the id:%d is %s\n", id, name)
		quant++
	}

	type user struct {
		Mensagem string `json:"msg"`
	}

	if quant == 0 {
		msg := fmt.Sprintf("Not Found for the id %d", id)
		fmt.Printf(msg, id)
		respondWithJson(responseWriter, 404, user{Mensagem: msg})
	}

	respondWithJson(responseWriter, 200, user{Mensagem: name})
}
