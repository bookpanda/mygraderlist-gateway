# MyGraderList Gateway

MyGraderList is a web app that lets students assess the difficulties and worthiness of each DSA grader problem in their respective courses.

MyGraderList Gateway handles the requests from the frontend and forwards them to the appropriate microservices.

## Technologies

-   golang
-   gRPC
-   go-fiber

## Getting Started

### Prerequisites

-   golang 1.21 or [later](https://go.dev)
-   docker
-   makefile

### Installation

1. Clone this repo
2. Copy `config.example.yaml` in 
- `config`
- `service/local/auth`
- `service/local/backend`


and paste it in the same directory with `.example` removed from its name. To use Google OAuth, you need to register a new credential in [Google Cloud Console](https://console.cloud.google.com/apis/dashboard)

3. Run `go mod download` to download all the dependencies.

### Running
1. Run `docker-compose up -d`
2. Run `make server` or `go run ./src/.`

### Testing
1. Run `make test` or `go test  -v -coverpkg ./... -coverprofile coverage.out -covermode count ./...`
