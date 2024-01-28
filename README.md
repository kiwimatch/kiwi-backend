# kiwimatch-backend

## Quick Start - Developer's Guide

> To install GO, visit https://go.dev/doc/install

Setup local environment
- `cp .env.local .env`

To run with kiwi-db(postgres) service
- Install docker then run `docker compose up --build` with -d flag to run in background
- `docker compose down` to stop all services

Running without db

- option 1: `go install github.com/cosmtrek/air@latest` then `air` (live reload)
- option 2: `go run .`

### Simple test for api working
`curl localhost:8080/api-health` should return with status 200

### Commit & Push
1. Run `go mod tidy` if you installed new packages with `go get` command
2. Do NOT push to main! Create new branch from main for all cases except documentations (exception permitted), then create a PR.
3. Keep it simple, no rebase necessary.