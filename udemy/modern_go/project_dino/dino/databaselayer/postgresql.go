package databaselayer
import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PQHandler struct {
	*SQLHandler
}

//return a type that implements the sqlhandler interface.
func NewPQHandler(connection string) (*PQHandler, error) {
	db, err := sql.Open("postgres", connection) // return db object as sqlhandler
	return &PQHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}