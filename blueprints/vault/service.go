package vault

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

// Service provides password hashing capabilities
type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate(ctx context.Context, password string, hash string) (bool, error)
}

type vaultService struct{}

//NewService makes a new Service. This prevents us from exposing our internals
//and lets us do more work to prepare vaultService without changing the API
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
