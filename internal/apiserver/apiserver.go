package apiserver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDb(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	srv := newServer()

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDb(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
