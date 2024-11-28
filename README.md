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
