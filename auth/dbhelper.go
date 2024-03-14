package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
