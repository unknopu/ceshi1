package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var(
	RepositoryCollection *DataStore
)

// MongoClient mongo client
type MongoClient interface {
	ping(ctx context.Context, rp *readpref.ReadPref) error
	database() *mongo.Database
}

type mongoClient struct {
	client *mongo.Client
}

// init new datastore
func NewDataStore(col *mongo.Collection) *DataStore {
	return &DataStore{collection: col}
}

func Init() {
	client := NewMongoClient()
	// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err := client.ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}

	RepositoryCollection = NewDataStore(client.database().Collection("repositories"))
	
	log.Println("Mongodb connected!! 🎉")
}

// NewMongoClient initial mongo connection
func NewMongoClient() MongoClient {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:2717")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// return client.Database(DB)
	return &mongoClient{
		client: client,
	}
}

func (c *mongoClient) ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.client.Ping(ctx, nil)
}

func (c *mongoClient) database() *mongo.Database {
	return c.client.Database("users")
}