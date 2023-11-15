package bd

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func inserirDados(db *sql.DB, nome string, idade int) error {
	// Consulta SQL para inserção
	insertSQL := "INSERT INTO arquivos (nome) VALUES ($1)"

	// Execute a consulta SQL para inserção
	_, err := db.Exec(insertSQL, nome, idade)
	return err
}
