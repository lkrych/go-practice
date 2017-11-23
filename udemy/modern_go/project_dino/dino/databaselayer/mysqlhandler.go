package databaselayer

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLHandler struct {
	*SQLHandler
}

//return a type that implements the sqlhandler interface.
func NewMySQLHandler(connection string) (*MySQLHandler, error) {
	db, err := sql.Open("postgres", connection) // return db object as sqlhandler
	return &MySQLHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}
