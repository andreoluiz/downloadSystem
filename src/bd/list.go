package bd

import (
	"database/sql"
	"fmt"
	"log"
)

func listAll(db *sql.DB, nome string, idade int) error {

	selectSQL := "SELECT * FROM arquivos"

	rows, err := db.Query(selectSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nome string
		if err := rows.Scan(&id, &nome); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Nome: %s\n", id, nome)
	}
	return err

}
