package models

import (
	"errors"
	"regexp"
	"strings"

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
	//ErrEmailRequired is returned when an email address is not provided when creating a user
	ErrEmailRequired = errors.New("models: email address is required")
	//ErrEmailInvalid is returned when an email address provided does not match our requirements
	ErrEmailInvalid = errors.New("models: email address is not valid")
	//ErrEmailTaken is returned when an update or create is attempted
	//with an email address that is already in use
	ErrEmailTaken = errors.New("models: email address is already taken")
	userPwPepper  = "thisisadummyvariable"
	hmacSecretKey = "this is another dummyVariable"
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

//userService is an abstraction layer that provides methods for querying, creating and updating users
type userService struct {
	UserDB //embedded interface satisfied by userGorm
}

//userGorm represents our database interaction layer
//and implements the UserDB interface fully.
type userGorm struct {
	db *gorm.DB
}

//userValidator is our validation layer that validates and
//nromalizes data before passing it on to the next
//userDB in our interface chain
type userValidator struct {
	UserDB
	hmac       hash.HMAC
	emailRegex *regexp.Regexp
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

// All of our validation functions need to accept a user pointer and return either
// nil or no errors
type userValFn func(*User) error

//UserService is a set of methods used to manipulate and work
//with the user model
type UserService interface {
	//Authenticate will verify the provided email address and
	//passwords are correct. If they are correct, the user corresponding
	//to that email will be returned. Otherwise you will receive either:
	//ErrNotFound, ErrInvalidPassword, or another error if something goes wrong
	Authenticate(email, password string) (*User, error)
	UserDB
}

//NewUserGorm instantiates a connection to the database in the userGorm object
func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &userGorm{
		db: db,
	}, nil
}

//NewUserValidator instantiates a userValidator object
func newUserValidator(udb UserDB, hmac hash.HMAC) *userValidator {
	return &userValidator{
		UserDB:     udb,
		hmac:       hmac,
		emailRegex: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
	}
}

//NewUserService returns an UserService interface
func NewUserService(connectionInfo string) (UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	hmac := hash.NewHMAC(hmacSecretKey)
	uv := newUserValidator(ug, hmac)
	return &userService{
		UserDB: uv,
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
	return ug.db.Create(user).Error
}

//Update will update the provided user with all the data in the provided user object
func (ug *userGorm) Update(user *User) error {
	return ug.db.Save(user).Error
}

//Delete will delete the user with the provided ID
func (ug *userGorm) Delete(id uint) error {
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
func (us *userService) Authenticate(email, password string) (*User, error) {
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
//and returns that the user. This method expects the remember token to be already hashed
func (ug *userGorm) ByRemember(rememberHash string) (*User, error) {
	var user User
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

//ByRemember will hash the remember token and then call
//ByRemember on the subsequent UserDB layer
func (uv *userValidator) ByRemember(token string) (*User, error) {
	user := User{
		Remember: token,
	}
	if err := runUserValFns(&user, uv.hmacRemember); err != nil {
		return nil, err
	}
	return uv.UserDB.ByRemember(user.RememberHash)
}

//ByEmail will normalize an email address before passing it on to
//the database layer to perform the query.
func (uv *userValidator) ByEmail(email string) (*User, error) {
	user := User{
		Email: email,
	}
	err := runUserValFns(&user, uv.normalizeEmail)
	if err != nil {
		return nil, err
	}
	return uv.UserDB.ByEmail(user.Email)
}

//Create will create the provided user and backfill data
//like the ID, CreatedAt, and UpdatedAt fields.
func (uv *userValidator) Create(user *User) error {

	err := runUserValFns(user,
		uv.bcryptPassword,
		uv.setRememberIfUnset,
		uv.hmacRemember,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emailIsAvail)
	if err != nil {
		return err
	}

	return uv.UserDB.Create(user)
}

//Update will hash a remember token if it is provided
func (uv *userValidator) Update(user *User) error {
	err := runUserValFns(user,
		uv.bcryptPassword,
		uv.hmacRemember,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emailIsAvail)
	if err != nil {
		return err
	}

	return uv.UserDB.Update(user)
}

//Delete will delete the user with the provided ID
func (uv *userValidator) Delete(id uint) error {
	var user User
	user.ID = id
	err := runUserValFns(&user, uv.idGreaterThan(0))
	if err != nil {
		return err
	}
	return uv.UserDB.Delete(id)
}

//bcryptPassword will hash a user's password with an
//app-wide pepper and bcrypt, which salts for us.
func (uv *userValidator) bcryptPassword(user *User) error {
	if user.Password == "" {
		//we DO NOT need to run this if the password
		//hasn't been changed
		return nil
	}
	pwBytes := []byte(user.Password + userPwPepper) //add pepper to password
	hashedBytes, err := bcrypt.GenerateFromPassword(
		pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//assign passwordFields
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return nil
}

//hmacRemember is a validator that hashes the rememberToken
func (uv *userValidator) hmacRemember(user *User) error {
	if user.Remember == "" {
		return nil
	}
	user.RememberHash = uv.hmac.Hash(user.Remember)
	return nil
}

//setRememberIfUnset will make sure tokens are given a default value
//if none exists
func (uv *userValidator) setRememberIfUnset(user *User) error {
	if user.Remember != "" {
		return nil
	}
	token, err := rand.RememberToken()
	if err != nil {
		return err
	}
	user.Remember = token
	return nil
}

//idGreaterThan uses a closure to dynamically create a validation function
// that verifies IDs must be greater than n
func (uv *userValidator) idGreaterThan(n uint) userValFn {
	return userValFn(func(user *User) error {
		if user.ID <= n {
			return ErrInvalidID
		}
		return nil
	})
}

//normalizeEmail trims the whitespace and changes all the characters to lowercase
func (uv *userValidator) normalizeEmail(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}

//requireEmail returns an error if an email hasn't been set
func (uv *userValidator) requireEmail(user *User) error {
	if user.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

//emailFormat checks to see if the email submitted matches our requirements
func (uv *userValidator) emailFormat(user *User) error {
	//we want to return the more valuable error from requireEmail, not this error
	if user.Email == "" {
		return nil
	}
	if !uv.emailRegex.MatchString(user.Email) {
		return ErrEmailInvalid
	}
	return nil
}

//emailIsAvail checks to see whether the submitted email has already been submitted
func (uv *userValidator) emailIsAvail(user *User) error {
	existing, err := uv.ByEmail(user.Email)
	if err == ErrNotFound {
		//email address is available
		return nil
	}
	//we can't continue our validation without a successful query,
	// so if we get any other error we should return it
	if err != nil {
		return err
	}

	//if we get here that means we found a user w/ this email
	// so we need to see if it is the same user we are updating,
	// or if we have a conflict
	if user.ID != existing.ID {
		return ErrEmailTaken
	}
	return nil
}

func runUserValFns(user *User, fns ...userValFn) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}
