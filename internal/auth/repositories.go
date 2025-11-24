package auth

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
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

func (repo *AuthRepository) FindUserByCredentials(email string, password string) (Credentials, error) {

	var credentials Credentials

	credentialsFilter := bson.D{{Key: "email", Value: email}, {Key: "password", Value: password}}

	err := repo.client.Database(repo.database).Collection(repo.collection).FindOne(context.TODO(), credentialsFilter).Decode(&credentials)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found")
		} else {
			fmt.Println("Something went wrong:", err)
		}

		return Credentials{}, err
	}

	fmt.Println(credentials.Email, credentials.Password)
	return credentials, nil
}
