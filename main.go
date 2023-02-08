package main

import (
	"api-golang/src/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config.Load()

	connectString := "postgres://xqnhqtbh:nfLvcjanlqbo13HzEDyNoUGr8x2LmQE2@babar.db.elephantsql.com/xqnhqtbh"

	db, erro := sql.Open("postgres", connectString)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexao aberta")
}
