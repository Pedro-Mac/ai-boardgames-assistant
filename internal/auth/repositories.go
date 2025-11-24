package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AuthRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewAuthRepository(client *mongo.Client, database string, collection string) *AuthRepository {
	return &AuthRepository{
		client:     client,
		database:   database,
		collection: collection,
	}
}

func (repo *AuthRepository) CreateUser(email string, password string) (*mongo.InsertOneResult, error) {
	authDoc := SignupRequestBody{
		Email:    email,
		Password: password,
	}

	return repo.client.Database(repo.database).Collection(repo.collection).InsertOne(context.TODO(), authDoc)

}
