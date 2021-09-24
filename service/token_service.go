package service

import (
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"context"
	"crypto/rsa"
	"log"
)

// TokenService used for injecting an implementation of TokenRepository
// for use in service methods along with keys and secret for
// signing JWTs
type TokenService struct {
	// TokenRepository model.TokenRepository
	PrivKey       *rsa.PrivateKey
	PubKey        *rsa.PublicKey
	RefreshSecret string
}

// TSConfig will hold repositories that will eventually be injected into this service layer
type TSConfig struct {
	// TokenRepository model.TokenRepository
	PrivKey       *rsa.PrivateKey
	PubKey        *rsa.PublicKey
	RefreshSecret string
}

// NewTokenService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewTokenService(c *TSConfig) model.TokenService {
	return &TokenService{
		PrivKey:       c.PrivKey,
		PubKey:        c.PubKey,
		RefreshSecret: c.RefreshSecret,
	}
}

// NewPairFromUser create fresh id and refresh tokens for the current user
// if a previous token is included, the previous token is remove
// from the tokens repository
func (s *TokenService) NewPairFromUser(ctx context.Context, u *model.User, prevToken string) (*model.TokenPair, error) {
	// No need to user a repository for idToken as it is unrelated to any data source
	idToken, err := generateIDToken(u, s.PrivKey)
	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v", u.UID, err.Error())
		return nil, apperrors.NewInternal()
	}

	refreshToken, err := generateRefreshToken(u.UID, s.RefreshSecret)
	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v", u.UID, err.Error())
		return nil, apperrors.NewInternal()
	}

	// TODO: store refresh tokens by calling TokenRepository method
	return &model.TokenPair{
		IDToken:      idToken,
		RefreshToken: refreshToken.SS,
	}, nil
}
