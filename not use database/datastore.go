package database

import (
	"ceshi1/account/model"
	"ceshi1/account/model/apperrors"
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)


type DataStore struct{
	collection *mongo.Collection
}

// datastore initilization
func NewDatastore(col *mongo.Collection) *DataStore {
	return &DataStore{collection: col}
}

// FindOne with query
func (ds DataStore) FindOne(ctx context.Context, filter interface{}) error {
	var u model.User
	return ds.collection.FindOne(ctx, filter).Decode(&u)
}

// Find by id
func (ds DataStore) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error){
	user := &model.User{}
	filter := bson.M{"uid": uid}

	err := ds.collection.FindOne(ctx, filter).Decode(user)
	if err != nil{
		return user, apperrors.NewNotFound("uid", uid.String())
	}
	return user, nil
}

