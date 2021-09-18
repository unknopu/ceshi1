package model

import (
	"context"

	"github.com/google/uuid"
)

// user service define method the handler layer expects
// any service it interacts with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
	// Create(ctx context.Context, u *User) error
}

// Token service defines methods the handler layer expects to internet
// with in regards to producing JWTs as string
type TokenService interface{
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
}

// user service define method the handler layer expects
// any service it interacts with to implement
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
	Create(ctx context.Context, u *User) error
}
