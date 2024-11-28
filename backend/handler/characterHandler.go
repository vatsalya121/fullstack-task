package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/vatsalya121/fullstack-task/backend/service"
    "github.com/vatsalya121/fullstack-task/backend/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// GetCharactersHandler fetches all characters from MongoDB
func GetCharactersHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		characters, err := service.FetchCharacters(client)
		if err != nil {
			http.Error(w, "Error fetching characters", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(characters)
	}
}

// GetCharacterByNameHandler fetches a character by name from MongoDB
func GetCharacterByNameHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		character, err := service.FetchCharacterByName(client, name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching character: %v", err), http.StatusInternalServerError)
			return
		}
		if character == nil {
			http.Error(w, "Character not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(character)
	}
}
