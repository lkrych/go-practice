package models

import (
	"errors"

	"flatphoto.com/hash"
	"flatphoto.com/rand"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	//ErrNotFound is returned when a resource cannot be found
	ErrNotFound = errors.New("models: resource not found")
	//ErrInvalidID is returned when invalid ID is provided to a method like delete
	ErrInvalidID = errors.New("models: ID provided was invalid")
	//ErrInvalidPassword is returned when an invalid password is used when attempting to authenticate a user
	ErrInvalidPassword = errors.New("models: incorrect password provided")
	userPwPepper       = "thisisadummyvariable"
	hmacSecretKey      = "this is another dummyVariable"
)

type User struct {
	gorm.Model
	Username     string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

//UserService is an abstraction layer that provides methods for querying, creating and updating users
type UserService struct {
	UserDB //embedded interface satisfied by userGorm
}

//UserDB is used to interact with the users database.
// For pretty much all single user queries:
// If the user is found we will return a nil error
// if the user is not found, we will return ErrNotFound
// if there is another error, we will return an error with
// more information about what went wrong
type UserDB interface {
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	//Methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	//Used to close a DB connection
	Close() error

	//Migration helpers
	AutoMigrate() error
	DestructiveReset() error
}

//userGorm represents our database interaction layer
//and implements the UserDB interface fully.
type userGorm struct {
	db   *gorm.DB
	hmac hash.HMAC
}

//userValidator is our validation layer that validates and
//nromalizes data before passing it on to the next
//userDB in our interface chain
type userValidator struct {
	UserDB
}

//NewUserGorm instantiates a connection to the database in the userGorm object
func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(hmacSecretKey)
	return &userGorm{
		db:   db,
		hmac: hmac,
	}, nil
}

//NewUserService instantiates a connection to the database in the Userservice object
func NewUserService(connectionInfo string) (*UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	return &UserService{
		UserDB: userValidator{
			UserDB: ug,
		},
	}, nil
}

//Close closes the UserService database connection
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

//ByID will look up a user with a provided ID
// If the user is found we will return a nil error
// if the user is not found, we will return ErrNotFound
// if there is another error, we will return an error with
// more information about what went wrong
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
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
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

//Create will create the provided user and backfill data like the ID, CreatedAt, etc.
func (ug *userGorm) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper) //add pepper to password
	hashedBytes, err := bcrypt.GenerateFromPassword(
		pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//assign passwordFields
	user.PasswordHash = string(hashedBytes)
	user.Password = ""

	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
	}
	user.RememberHash = ug.hmac.Hash(user.Remember)
	return ug.db.Create(user).Error
}

//Update will update the provided user with all the data in the provided user object
func (ug *userGorm) Update(user *User) error {
	if user.Remember != "" {
		user.RememberHash = ug.hmac.Hash(user.Remember)
	}
	return ug.db.Save(user).Error
}

//Delete will delete the user with the provided ID
func (ug *userGorm) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

//DestructiveReset drops the user table and rebuilds it
//only useful for development
func (ug *userGorm) DestructiveReset() error {
	err := ug.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return ug.AutoMigrate()

}

//AutoMigrate will attempt to automatically migrate the users table
func (ug *userGorm) AutoMigrate() error {
	if err := ug.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

//Authenticate can be used to authenticate a user with the provided email and password
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash),
		[]byte(password+userPwPepper),
	)
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrInvalidPassword
	default:
		return nil, err
	}
}

//ByRemember looks up a user with the given remember token
//and returns that the user. This method will handle hashing the token for us
func (ug *userGorm) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := ug.hmac.Hash(token)
	err := first(ug.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
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
