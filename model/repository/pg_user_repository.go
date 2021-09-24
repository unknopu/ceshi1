package repository

import (
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

// PGUserRepository is data/repository implementation of service layer UserRepository
type pgUserRepository struct {
	DB *sqlx.DB
}

// the factory for initializing User repositories
func NewUserRepository(db *sqlx.DB) model.UserRepository {
	return &pgUserRepository{
		DB: db,
	}
}

// Create reaches out to database SQLX api
func (r *pgUserRepository) Create(ctx context.Context, u *model.User) error{
	query := "insert into users (email, password) values ($1, $2) returning *"

	if err := r.DB.Get(u, query, u.Email, u.Password); err != nil {
		// check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a user with email: %v. Reason: %v", u.Email, err.Code.Name())
			return apperrors.NewConflict("emiail", u.Email)
			
		}

		log.Printf("Could not create a user with email: %v. Reason: %v", u.Email, err)
		return apperrors.NewInternal()
	}
	return nil
}
