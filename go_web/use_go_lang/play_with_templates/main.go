package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "flatphoto_dev"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm: not null;unique_index`
	Orders []Order
}

//declaring order type
type Order struct {
	gorm.Model
	UserId      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// err = db.Ping() //make sure that our code actually talks to the db
	// if err != nil {
	// 	panic(err)
	// }

	//the following three examples make use of the import statement
	// _ "github.com/lib/pq"

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

	// var id int
	// for i := 1; i < 6; i++ {
	// 	//Create some fake data
	// 	userId := 1
	// 	if i > 3 {
	// 		userId = 3
	// 	}
	// 	amount := 1000 * i
	// 	description := fmt.Sprintf("USB-C Adapter x%d", i)

	// 	err = db.QueryRow(`
	// 		INSERT INTO orders(user_id, amount, description)
	// 		VALUES ($1, $2, $3)
	// 		RETURNING id`, userId, amount, description).Scan(&id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Created an order with id:", id)
	// }

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//enable db logging
	db.LogMode(true)

	//takes an instance of several types and creates corresponding tables
	//it can also be used to update models
	//it will only create things that don't exist
	db.AutoMigrate(&User{}, &Order{})

	//creating a record with GORM
	// name, email := getInfo()
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }
	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", u)

	// // grab the first user
	var firstUser User
	db.Preload("Orders").First(&firstUser)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println("The first user is ", firstUser.Name)
	fmt.Println("They have ordered", len(firstUser.Orders), "items")
	fmt.Println("Orders:", firstUser.Orders)

	//  grab a user by id
	// var secondUser User
	// db.First(&secondUser, 2)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }
	// fmt.Println("The second user is ", secondUser)

	// // chaining Where
	// var whereUser User
	// maxId := 3

	// db.Where("id <= ?", maxId).First(&whereUser)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }
	// fmt.Println("Using where: ", whereUser)

	//grabbing multiple users
	// var users []User
	// db.Find(&users)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }
	// fmt.Println("Retrieved", len(users), "users.")
	// fmt.Println(users)

	//adding orders
	// createOrder(db, firstUser, 1001, "Fake Description #1")
	// createOrder(db, firstUser, 9999, "Fake Description #2")
	// createOrder(db, firstUser, 2, "Fake Description #3")

	fmt.Println("Successfully connected to the db!")
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("what is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserId:      user.ID,
		Amount:      amount,
		Description: desc,
	})
	if db.Error != nil {
		panic(db.Error)
	}
}
