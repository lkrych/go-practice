package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "root:@/dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//query with arguments
	rows, err := db.Query("select * from dino.animals where age > ?", 9)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age) // mysql driver return data for row
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

	//query a single row
	row := db.QueryRow("select * from dino.animals where age > ?", 10)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age) // mysql driver return data for row
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	//insert a row

	// result, err := db.Exec("Insert into dino.animals(animal_type, nickname, zone, age) values ('Butterfly', 'Ludwig', 3, 1)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())

	//update a row

	result, err := db.Exec("Update dino.animals set age = ? where id = ?", 16, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}
