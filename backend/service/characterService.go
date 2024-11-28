package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"backend/model"
)

// ConnectToMongoDB connects to the MongoDB instance and returns a client.
func ConnectToMongoDB() (*mongo.Client, error) {
	// MongoDB connection string
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://myuser:1234@35.232.163.34:27017/rick_and_morty"))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	// Attempt to connect
	err = client.Connect(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Optionally check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("MongoDB ping failed: %v", err)
	}

	// If everything is okay, return the client
	return client, nil
}

// FetchCharacters fetches all characters from the MongoDB collection.
func FetchCharacters(client *mongo.Client) ([]model.Character, error) {
	collection := client.Database("rick_and_morty").Collection("characters")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find characters: %v", err)
	}
	defer cursor.Close(context.Background())

	var characters []model.Character
	for cursor.Next(context.Background()) {
		var character model.Character
		err := cursor.Decode(&character)
		if err != nil {
			return nil, fmt.Errorf("failed to decode character: %v", err)
		}
		characters = append(characters, character)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %v", err)
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
			return nil, fmt.Errorf("character not found")
		}
		return nil, fmt.Errorf("failed to find character by name: %v", err)
	}
	return &character, nil
}
