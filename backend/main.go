package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/vatsalya121/fullstack-task/backend/handler"
    "github.com/vatsalya121/fullstack-task/backend/service"
    "github.com/vatsalya121/fullstack-task/backend/model"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize MongoDB connection
	client, err := service.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(nil)
	
	// Initialize router
	router := mux.NewRouter()

	// Set up routes
	router.HandleFunc("/api/characters", handler.GetCharactersHandler(client)).Methods("GET")
	router.HandleFunc("/api/characters/{name}", handler.GetCharacterByNameHandler(client)).Methods("GET")

	// Start server
	fmt.Println("Server started on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
