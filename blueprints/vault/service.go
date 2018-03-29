package vault

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Service provides password hasing capabilities.
type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate(ctx context.Context, password string, hash string) (bool, error)
}

type vaultService struct{}

type hashRequest struct {
	Password string `json:"password"`
}
type hashResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

type validateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}
type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

//NewService makes a new Service!
func NewService() Service {
	return vaultService{}
}

func (vaultService) Hash(ctx context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (vaultService) Validate(ctx context.Context, password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

//user -> server *we are here* -> rpc -> server -> user
//decodeHashRequest is a helper method that decodes the body of of a JSON http request into the struct we need for our service
func decodeHashRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req hashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//user -> server *we are here* -> rpc -> server -> user
//decodeValidateRequest is a helper method that decodes the body of of a JSON http request into the struct we need for our service
func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//user -> server -> rpc -> server *we are here* -> user

//encodeResponse can handle both hashResponse and validateResponse
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
