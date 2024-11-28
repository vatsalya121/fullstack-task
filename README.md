# Fullstack Task Solution

This repository contains the solution for the Fullstack Task, which includes three main parts:

1. Cloud Infrastructure (GCP + Terraform)
2. Backend (Go)
3. Frontend (React + TypeScript)

## Part 1: Cloud Infrastructure (GCP + Terraform)

### Requirements:
- Created Terraform configuration for MongoDB Community Edition deployment on Google Cloud Platform.
- Configured Google Compute Engine and attached Block storage.
- Set up firewall rules to protect the instance.
- Downloaded and stored all characters from the Rick and Morty API.
- Destroyed the infrastructure and restored it from block storage.

### Key Features:
- Clean and readable Terraform code.
- MongoDB Community Edition hosted on Google Cloud Platform.
- Compute Engine and Block Storage configured using Terraform.
- Proper backup and restore functionality.

### Setup Instructions:
1. Install Terraform and Google Cloud SDK.
2. Configure Google Cloud credentials (`gcloud auth login`).
3. Run the following Terraform commands:
   ```bash
   terraform init
   terraform apply
4.To destroy the infrastructure:
```bash
terraform destroy
```
##Part 2: Backend (Go)

### Requirements:
-Built a JSON REST API using Go to connect to MongoDB.
-Implemented search functionality by character name to fetch character details.

### Key Features:
-Clean Go code architecture.
-Error handling implemented.
-MongoDB connection and search functionality.

### Setup Instructions:
1. Clone the repository.
2. Install dependencies:
```bash
go mod tidy
```
3. Run the Go backend:
```bash
go run main.go
```
4. The API will be accessible at http://localhost:8080.
   
### Endpoints:
-GET /characters - Fetch all characters.
-GET /characters/{name} - Fetch character details by name.

## Part 3: Frontend (React + TypeScript)

### Requirements:

-Built a search engine for character names.
-Displayed character details and associated episode numbers.
-Handled loading and error states.
-Added basic styling using TailwindCSS.

### Key Features:

-TypeScript implementation.
-Component structure and state management using React.
-TailwindCSS for styling.

### Setup Instructions:

1. Clone the repository.
2. Navigate to the frontend folder:
   ```bash
   cd frontend
   ```
3. Install dependencies:
   ```bash
   npm install
   ```
4. Run the React app:
```bash
   npm start
```
5. The frontend will be available at http://localhost:3000.

## Components:
-SearchBar - Allows the user to search for characters by name.
-CharacterCard - Displays character details and associated episodes.

## Styling:
-Basic styling is done using TailwindCSS for responsiveness and clean UI design.

## Screenshot of Frontend UI:

### Architecture Overview

## MongoDB Setup (Cloud Infrastructure):
-MongoDB is deployed on a Google Compute Engine instance with Block Storage.
-The database stores character data fetched from the Rick and Morty API.

## Backend (Go):
-The Go backend exposes REST APIs to interact with MongoDB.
-The backend connects to MongoDB and provides the search functionality for character names.

## Frontend (React + TypeScript):
-The frontend is a React application with TypeScript.
-It communicates with the backend to fetch character data and display it to the user.

## Testing Instructions

# Backend:

-Ensure MongoDB is running locally or on the cloud.
-Run the backend using go run main.go.
-Test the API endpoints using Postman or curl.
-GET /characters - Fetch all characters.
-GET /characters/{name} - Fetch character details by name.

# Frontend:
-Ensure the backend is running locally at http://localhost:8080.
-Run the frontend with npm start and test the UI.
-Test the character search functionality and ensure proper error handling.

### Directories Structure
```bash
/terraform
    ├── main.tf

/Backend
    ├── main.go
    ├── handler
    └── service

/Frontend
    ├── src
    ├── components
    ├── images
    └── tailwind.config.js
```
### GitHub Setup

-Create a GitHub repository and push the code to the repository.
-Use pull requests for code merging and squash commits while merging.
-Do not force push and avoid forking the repository.
-Ensure that the README.md contains all the setup and documentation details.




