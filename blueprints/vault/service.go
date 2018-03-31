package vault

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
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

//MakeHashEndpoint is a function that handles an incoming hashRequest, calls the Hash Service method, and returns the appropriate hashResponse object
//Endpoints are special functions that represent a single RPC method in GoKit
func MakeHashEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hashRequest)
		v, err := srv.Hash(ctx, req.Password)
		if err != nil {
			return hashResponse{v, err.Error()}, nil
		}
		return hashResponse{v, ""}, nil
	}
}

//MakeValidateEndpoint is a function that handles an incoming validateRequest, calls the Validate Service method, and returns the appropriate validateResponse object
func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return validateResponse{false, err.Error()}, nil
		}
		return validateResponse{v, ""}, nil
	}
}

//An implementation of the vault.Service interface
type Endpoints struct {
	HashEndpoint     endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

//Implementing the hash method for the endpoints struct
//Notice that there are different levels of errors, top level (transport) errors and business logic errors
func (e Endpoints) Hash(ctx context.Context, password string) (string, error) {
	req := hashRequest{Password: password}
	resp, err := e.HashEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	hashResp := resp.(hashResponse)
	if hashResp.Err != "" {
		return "", errors.New(hashResp.Err)
	}
	return hashResp.Hash, nil
}

//Implementing the validate method for the endpoints struct,
//*these two methods will allow us to treat the endpoints as if they were normal go methods
func (e Endpoints) Validate(ctx context.Context, password string, hash string) (bool, error) {
	req := validateRequest{Password: password, Hash: hash}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(validateResponse)
	if validateResp.Err != "" {
		return false, errors.New(validateResp.Err)
	}
	return validateResp.Valid, nil
}
