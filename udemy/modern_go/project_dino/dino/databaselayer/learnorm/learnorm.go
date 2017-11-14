package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type animal struct {
	gorm.Model
	AnimalType string `gorm:"type:TEXT"`
	Nickname   string `gorm:"type:TEXT"`
	Zone       int    `gorm:"type:INTEGER"`
	Age        int    `gorm:"type:INTEGER"`
}

func main() {
	db, err := gorm.Open("sqlite3", "dino.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.DropTableIfExists(&animal{})

	//checks if table exists, if not, it creates the table and columns of the struct type
	db.AutoMigrate(&animal{}) // will add any missing fields, will add 's' to the struct name
	// db.Table("dinos").CreateTable(&animal{}) //create a custom tablename

	a := animal{
		AnimalType: "Beluga Whale",
		Nickname:   "Whitey",
		Zone:       1,
		Age:        5,
	}
	db.Save(&a) // will create a new row if no ID is given

	a = animal{
		AnimalType: "Orangutan",
		Nickname:   "Orange Crush",
		Zone:       2,
		Age:        3,
	}
	db.Save(&a)

	//updates

	// db.Table("animals").Where("Nickname = ?", "Orange Crush").Update("Nickname", "BANANA SLUG")

	//queries
	animals := []animal{}
	//Debug shows what SQL commands are being run by gorm
	// db.Debug().Find(&animals, "age > ?", 2)
	db.Find(&animals, "age > ?", 2)
	fmt.Println(animals)

}
