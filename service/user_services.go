package service

import (
	"ceshi1/account/model"
	"context"

	"github.com/google/uuid"
)

// UserService acts as a struct for injecting an implementation for UserRepository
// for use in service methods
type UserService struct{
	UserRepository model.UserRepository
}

// USConfig will hold repositories that will eventually be injected into this
// this service layer
type USConfig struct {
	UserRepository model.UserRepository	
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService{
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

func (s *UserService) Get(ctx context.Context,uid uuid.UUID) (*model.User, error){
	u, err := s.UserRepository.FindByID(ctx, uid)
	return u, err
}


// signup reacher our to a UserRepository to verify that,
// the email address is available and signs up the user if this is the case
func (s *UserService) Signup(ctx context.Context, u *model.User) error{
	panic("Method not implement yet.")
}

