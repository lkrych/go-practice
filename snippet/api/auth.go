package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//Struct definitions
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

const (
	privKeyPath = "./api/keys/app.rsa"
	pubKeyPath  = "./api/keys/app.rsa.pub"
)

var VerifyKey, SignKey []byte

func InitKeys() {

	var err error
	//These keys sign your token
	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("Error reading private key: ", err)
		return
	}
	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key: ", err)
		return
	}
}

//LoginHandler is used to create and sign JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user UserCredentials

	//decode request into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Fatalf("Error in request: %v\n", err)
		fmt.Fprintf(w, "Error in request")
	}

	fmt.Println(user.Username, user.Password)

	//validate user credentials
	if strings.ToLower(user.Username) != "noodle" {
		if user.Password != "bop" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error loggin in")
			fmt.Fprintf(w, "Invalid credentials")
		}
	}

	//create a rsa256 signer
	signer := jwt.New(jwt.GetSigningMethod("RS256"))

	//set claims
	//set claims
	claims := make(jwt.MapClaims)
	claims["iss"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24)
	claims["CustomUserInfo"] = struct {
		Name string
		Role string
	}{user.Username, "Member"}
	signer.Claims = claims

	tokenString, err := signer.SignedString(SignKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing token")
		log.Printf("Error signing token: %v\n", err)
	}

	//create a token instance using the token string
	response := Token{tokenString}
	JsonResponse(response, w)

}

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//validate token
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil
	})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}
