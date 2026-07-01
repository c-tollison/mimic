# Mimic

CLI for running Postgres migrations, with built-in database branching so you can spin up isolated workspaces instead of resetting your local DB.

## Build & Run

```bash
# Build the binary
go build -o mimic .

# Run directly without building
go run .

# Install binary to your Go bin, lets you run `mimic` from anywhere
go install .
```

---

## Project Structure

```
mimic/
├── main.go           # Entry point — calls cmd.Execute()
├── go.mod
├── go.sum
├── cmd/
│   └── root.go       # Root Cobra command
└── internal/
    └── db/
        └── client.go
```
