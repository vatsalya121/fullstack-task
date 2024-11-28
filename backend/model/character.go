package model

// Character represents the structure of a character in the database.
type Character struct {
	ID      string `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Status  string `json:"status" bson:"status"`
	Species string `json:"species" bson:"species"`
	Gender  string `json:"gender" bson:"gender"`
}
