package repository

import (
	"ceshi1/account/database"
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// MGUserRepository is data/repository implementation
// of service layer UserRepository
type MGUserRepository struct {
	DB *mongo.Database
}



// create reaches out to database mongo api
func (r *MGUserRepository) Create(ctx context.Context, u *model.User) error {
	filter := bson.M{"email": u.Email}
	err := database.RepositoryCollection.FindOne(ctx, filter)



	if err != nil {
		log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err.Error)
			return apperrors.NewConflict("email", u.Email)
	}
	fmt.Println(err)

	return nil
}
