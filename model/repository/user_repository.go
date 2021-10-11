package repository

import (
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PGUserRepository is data/repository implementation of service layer UserRepository
type UserRepository struct {
	DB *gorm.DB
}

// the factory for initializing User repositories
func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// Create reaches out to database SQLX api
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	var user model.User
	r.DB.Where("email = ?", u.Email).Find(&user)
	
	if user.Email != u.Email {
		user.UID = uuid.New()
		user.Email = u.Email
		user.Password = u.Password
		r.DB.Create(user)
		return nil
	}

	return apperrors.NewConflict("email", u.Email)
}

// FindByID fetches user by id
func (r *UserRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}

	// query := "SELECT * FROM users WHERE uid=$1"
	// if err := r.DB.GetContext(ctx, user, query, uid); err != nil {
	// 	return user, apperrors.NewNotFound("uid", uid.String())
	// }

	// we need to actually check errors as it could be something other than not found
	if err := r.DB.Table("users").Where("user_name=?", uid).Error; err != nil {
		return user, apperrors.NewNotFound("uid", uid.String())
	}

	return user, nil
}
