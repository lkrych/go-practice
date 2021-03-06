package databaselayer

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PGHandler struct {
	*SQLHandler
}

//return a type that implements the sqlhandler interface.
func NewPGHandler(connection string) (*PGHandler, error) {
	db, err := sql.Open("postgres", connection) // return db object as sqlhandler
	return &PGHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
