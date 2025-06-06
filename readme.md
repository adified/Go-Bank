# GoBankAPI 🏦

A **production-grade banking backend** built in **Golang** using the Gin framework. It supports secure user authentication, account management, fund transfers, transaction history, and exposes both **RESTful** and **gRPC** APIs. The project follows clean architecture principles and includes containerization, CI/CD, and cloud-native deployment.

---

## 🚀 Features

- ✅ REST and gRPC API support via `gRPC-Gateway`
- 🔐 JWT + Paseto authentication with refresh tokens
- 🛂 Role-Based Access Control (RBAC)
- 💸 Secure fund transfers with database transactions
- 🗃️ PostgreSQL with `pgx` connection pooling
- 🔁 Asynchronous task processing using Redis + Asynq
- 🧾 Structured API documentation with Swagger
- 📦 Dockerized and deployed using AWS (ECR + EKS)
- 🔄 CI/CD with GitHub Actions
- 📊 Observability using Prometheus + Grafana (Planned)
- 📂 Migrations managed with `migrate`
- 🧪 90%+ unit test coverage with `Testify` and `GoMock`

---

## 🧱 Tech Stack

| Component        | Technology                        |
| ---------------- | --------------------------------- |
| Language         | Golang                            |
| Framework        | Gin                               |
| Auth             | JWT, Paseto                       |
| DB               | PostgreSQL, pgx                   |
| Migrations       | golang-migrate                    |
| Async Queue      | Redis + Asynq                     |
| API              | REST, gRPC, gRPC-Gateway, Swagger |
| Testing          | Testify, GoMock                   |
| CI/CD            | GitHub Actions                    |
| Containerization | Docker                            |
| Deployment       | AWS ECR, ECS, EKS (Kubernetes)    |
| Monitoring       | Prometheus, Grafana (planned)     |
| Logging          | Zerolog                           |

---

## 📂 Folder Structure

├── api/ # REST API layer (HTTP handlers, routing)\
├── db/ # Database-related code\
│ ├── migrate/ # SQL migration files\
│ ├── migration/ # Migration utility logic\
│ ├── mock/ # Mock implementations for tests\
│ ├── query/ # Manually written SQL queries (if any)\
│ ├── sqlc/ # Auto-generated SQLC Go code\
├── doc/ # Documentation files\
├── eks/ # Kubernetes/EKS deployment configuration\
├── gapi/ # gRPC API server\
├── mail/ # Email service logic (SMTP etc.)\
├── pb/ # Protocol buffer compiled files (Go stubs)\
├── proto/ # gRPC proto definitions\
├── token/ # Token creation and validation logic (JWT/PASETO)\
├── util/ # Utility/helper functions\
├── val/ # Input validation logic\
├── worker/ # Background job processing (e.g., with Redis or Task queues)\

---

## Root-level files

├── .env # Environment variables for local dev\
├── .gitignore # Git ignore file\
├── app.env # App-specific environment config\
├── docker-compose.yml # Docker services config (DB, API, etc.)\
├── Dockerfile # Go application container\
├── go.mod # Go modules definition\
├── go.sum # Go modules checksums\
├── main.go # Main entry point of the application\
├── Makefile # Automation scripts for build/test/migrate\
├── readme.md # Project overview (this file)\
├── sqlc.yaml # SQLC configuration file\
├── start.sh # Startup script (optional)\

---

## ⚙️ Setup Instructions

### Prerequisites

- Go 1.20+
- PostgreSQL
- Docker + Docker Compose
- Redis
- Protobuf Compiler (`protoc`)
- `sqlc`, `migrate`

```bash
git clone https://github.com/adified/Go-Bank.git
cd Go-Bank

make docker-up   # Runs Postgres, Redis, etc. via Docker Compose
make migrate-up  # Runs DB migrations
make server      # Starts the backend API server
make test        # Runs unit tests
make mockgen     # Generates mocks using GoMock
```

### Swagger available at: http://localhost:8080/swagger/index.html

### gRPC services defined in api/proto/\*.proto

---

## 📦 Deployment

The app is production-ready with Docker and Kubernetes manifests.

Container images are pushed to AWS ECR.

Deployed on AWS EKS using.

CI/CD automated with GitHub Actions.
