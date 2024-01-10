# Snip-Snap

A pastebin clone written in Go

# Setup

- Make sure go is installed, and install all dependencies
- Generate TLS certificates by following the instruction inside the `/tls` folder
- Setup a PostgreSQL server, and run the SQL files inside `./migrations`
- Start the program using by `go run ./cmd/web --dsn=<postgresql_connection_string>`
