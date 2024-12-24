Full-Stack Web Application with CI/CD Pipeline
Overview
This repository contains a full-stack web application deployed with a robust CI/CD pipeline using GitHub Actions, Google Cloud Platform (GCP), and Turso database.
Infrastructure
Continuous Integration
Our CI pipeline runs on every Pull Request and ensures code quality through:

✅ Unit Testing
✅ Code Formatting Checks
✅ Linting
✅ Security Scans

Continuous Deployment
When changes are merged into the main branch, our CD pipeline automatically:

Builds the server binary
Creates a Docker image
Pushes the image to Google Artifact Registry
Deploys to Google Cloud Run with zero-downtime deployment

Database

Powered by Turso, a cloud-hosted SQLite database
Ensures reliable data persistence
Managed database solution

Technology Stack

Cloud Platform: Google Cloud Platform (GCP)

Cloud Run for serverless container deployment
Artifact Registry for Docker image storage


Database: Turso (Cloud-hosted SQLite)
CI/CD: GitHub Actions
Containerization: Docker

Infrastructure Features

Automated deployment process
Zero-downtime deployments
Automatic scaling
SSL/TLS certificate management
Load balancing
DNS management

Development Notes
Local Development
[Add instructions for local development setup]
Deployment
The application automatically deploys when changes are merged into the main branch. No manual deployment steps are required.


![CI](https://github.com/Sureshocked/learn-cicd-starter/actions/workflows/ci.yml/badge.svg)
# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!
Sureshocked's version of Boot.dev's Notely app.
Sureshocked's version of Boot.dev's Notely app.
Additional line of content at the bottom of the README.
