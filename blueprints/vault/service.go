package vault

import "context"

// Service provides password hasing capabilities.
type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate(ctx context.Context, password string, hash string) (bool, error)
}

type vaultService struct{}
