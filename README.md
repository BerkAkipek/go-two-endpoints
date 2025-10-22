# Go REST API with Docker and MySQL

A minimal yet production-ready Go REST API built with the [Chi router](https://github.com/go-chi/chi), using MySQL as a database, Docker Compose for orchestration, and environment-based configuration for development and production.

---

## Project Structure

go-two-endpoints/
├── Dockerfile
├── docker-compose.yml
├── .env.dev
├── .env.prod
├── go.mod
├── go.sum
├── main.go
├── routes/
│ └── routes.go
└── handlers/
  └── handlers.go


---

## Features

- Written in Go (v1.25 LTS)
- RESTful API using Chi Router
- MySQL database with Docker volume persistence
- Secure production setup with Docker Secrets
- Environment-based configuration (.env.dev / .env.prod)
- Multi-stage Dockerfile for small, non-root production images

---

## Endpoints

| Method | Endpoint        | Description        |
|:-------|:----------------|:-------------------|
| GET    | `/api/users`    | Returns all users  |
| POST   | `/api/users`    | Creates a new user |

### Example

**POST** `/api/users`

Request:
```json
{
  "name": "Alice",
  "email": "alice@example.com"
}
```
Response:
```json
{
  "message": "User created",
  "data": {
    "name": "Alice",
    "email": "alice@example.com"
  }
}
```

Running with Docker Compose

Development
docker compose --env-file .env.dev up --build


The application runs on http://localhost:8080

Production
docker compose --env-file .env.prod up --build -d


The application runs on http://localhost

Secure Secrets (Production)

Database credentials and sensitive data are stored in Docker secrets (./secrets/ directory), not in .env.prod.

Example structure:

secrets/
├── db_user.txt
├── db_password.txt
└── db_name.txt

Tech Stack

Language: Go 1.25 LTS

Router: go-chi/chi

Database: MySQL 8

Containerization: Docker & Docker Compose

Secrets Management: Docker Secrets

Environment Config: .env files

Common Commands
Action	Command
Build project	go build -o main .
Run locally (Go only)	go run main.go
Docker build	docker build -t go-two-endpoints .
Compose up	docker compose up --build
Compose down	docker compose down -v
Example .env.dev
DB_HOST=db
DB_USER=root
DB_PASSWORD=secret
DB_NAME=goapp_dev
APP_PORT=8080

Author

Berk Akipek
Backend and Cloud Engineer