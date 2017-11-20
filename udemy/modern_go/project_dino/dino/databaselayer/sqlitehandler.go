package databaselayer
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteHandler struct {
	*SQLHandler
}

//return a type that implements the sqlhandler interface.
func NewSQLiteHandler(connection string) (*SQLiteHandler, error) {
	db, err := sql.Open("sqlite3", connection)
	return &SQLiteHandler{
		SQLHandler: &SQLHandler{
			DB: db,
		},
	}, err
}