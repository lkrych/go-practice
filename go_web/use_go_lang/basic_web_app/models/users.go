package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	//ErrNotFound is returned when a resource cannot be found
	ErrNotFound = errors.New("models: resource not found")
	//ErrInvalidID is returned when invalid ID is provided to a method like delete
	ErrInvalidID = errors.New("models: ID provided was invalid")
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

//abstraction layer that provides methods for querying, creating and updating users
type UserService struct {
	db *gorm.DB
}

//NewUserService instantiates a connection to the database in the Userservice object
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

//Close closes the UserService database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

//ByID will look up a user with a provided ID
// If the user is found we will return a nil error
// if the user is not found, we will return ErrNotFound
// if there is another error, we will return an error with
// more information about what went wrong
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//ByEmail looks up a user with the given email address and
// returns that user.
// If the user is found we will return a nil error
// if the user is not found, we will return ErrNotFound
// if there is another error, we will return an error with
// more information about what went wrong
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

//Create will create the provided user and backfill data like the ID, CreatedAt, etc.
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

//Update will update the provided user with all the data in the provided user object
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

//Delete will delete the user with the provided ID
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

//DestructiveReset drops the user table and rebuilds it
//only useful for development
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

// first will query using the provided gorm.Db and it will
// get the first item returned and place it into dst. If
// nothing is found it he query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}