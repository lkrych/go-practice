package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "flatphoto"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping() //make sure that our code actually talks to the db
	if err != nil {
		panic(err)
	}

	//inserting a record
	// var id int
	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2) RETURNING id`,
	// 	"Spaghetti", "Alfredo@sauce.com")
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// grabbing a particular row by id
	// var id int
	// var name, email string
	// row := db.QueryRow(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE ID=$1`, 1)
	// err = row.Scan(&id, &name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("ID:", id, "Name:", name, "Email:", email)

	// doing a specific SQL query
	// var id int
	// var name, email string
	// rows, err := db.Query(`
	// 	SELECT id, name, email
	// 	FROM users
	// 	WHERE email = $1
	// 	OR ID > $2`,
	// 	"noodlebop@gmail.com", 3)
	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	rows.Scan(&id, &name, &email)
	// 	fmt.Println("ID:", id, "Name:", name, "Email:", email)
	// }

	var id int
	for i := 1; i < 6; i++ {
		//Create some fake data
		userId := 1
		if i > 3 {
			userId = 3
		}
		amount := 1000 * i
		description := fmt.Sprintf("USB-C Adapter x%d", i)

		err = db.QueryRow(`
			INSERT INTO orders(user_id, amount, description)
			VALUES ($1, $2, $3)
			RETURNING id`, userId, amount, description).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Created an order with id:", id)
	}
	fmt.Println("Successfully connected to the db!")
	db.Close()
}
