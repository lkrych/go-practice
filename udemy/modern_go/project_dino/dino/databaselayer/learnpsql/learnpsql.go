package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	//connect to the database
	db, err := sql.Open("postgres", "user=postgres dbname=dino sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// //general query with arguments
	// rows, err := db.Query("select * from animals where age > $1", 9)
	// handlerows(rows, err)

	// //query a single row
	// row := db.QueryRow("select * from animals where age > $1", 10)
	// a := animal{}
	// err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age) // psql driver return data for row
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(a)

	//insert a row

	// result, err := db.Exec("Insert into animals(animal_type, nickname, zone, age) values ('Butterfly', 'Ludwig', 3, 1)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.LastInsertId()) //not supported in psql
	// fmt.Println(result.RowsAffected())

	//update a row

	// result, err := db.Exec("Update animals set age = $1 where id = $2", 16, 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.LastInsertId()) // not supported in psql
	// fmt.Println(result.RowsAffected())

	//retrieve last effected id
	// var id int
	// db.QueryRow("Update animals set age = $1 where id = $2", 16, 2).Scan(&id)
	// fmt.Println("id returned:", id)

	//prepare queries to use them multiple times, saved as index, more efficient than interpreting sql every time
	// fmt.Println("Statements... ")
	// stmt, err := db.Prepare(" select * from animals where age > $1 ")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// //let's try with age > 0
	// rows, err := stmt.Query(0)
	// handlerows(rows, err)

	// //lets try with age > 10
	// rows, err = stmt.Query(10)
	// handlerows(rows, err)

	testTransaction(db)
}

func handlerows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age) // psql driver return data for row
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	if err := rows.Err(); err != nil { //make sure errors encountered during iteration are logged
		log.Fatal(err)
	}
	fmt.Println(animals)
}

func testTransaction(db *sql.DB) {
	//transactions are isolated to a single connection in the DB
	//by contrast a db handler sits on top of a connection pool
	fmt.Println("Transactions...")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() //rolls back any uncommitted SQL commands

	stmt, err := tx.Prepare(" select * from animals where age > $1 ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(10)
	handlerows(rows, err)
	rows, err = stmt.Query(14)
	handlerows(rows, err)

	results, err := tx.Exec("Update animals set age = $1 where id = $2", 11, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.RowsAffected())

	err = tx.Commit() //commits transaction into DB
	if err != nil {
		log.Fatal(err)
	}
}
