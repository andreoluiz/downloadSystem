package bd

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectAndPing() {
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else if err == nil {
		fmt.Println("Conexão com o PostgreSQL bem-sucedida!")
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS arquivos (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(255) UNIQUE,
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}


}
