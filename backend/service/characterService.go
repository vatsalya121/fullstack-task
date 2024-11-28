package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/vatsalya121/fullstack-task/backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB connects to the MongoDB instance and returns a client.
func ConnectToMongoDB() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://myuser:1234@35.232.163.34:27017/rick_and_morty"))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// FetchCharacters fetches all characters from the MongoDB collection.
func FetchCharacters(client *mongo.Client) ([]model.Character, error) {
	collection := client.Database("rick_and_morty").Collection("characters")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var characters []model.Character
	for cursor.Next(context.Background()) {
		var character model.Character
		err := cursor.Decode(&character)
		if err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}
	return characters, nil
}

// FetchCharacterByName fetches a character by name from the MongoDB collection.
func FetchCharacterByName(client *mongo.Client, name string) (*model.Character, error) {
	collection := client.Database("rick_and_morty").Collection("characters")
	var character model.Character
	err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&character)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Character not found
		}
		return nil, err
	}
	return &character, nil
}
