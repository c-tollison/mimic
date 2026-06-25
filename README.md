# Mimic

Postgres migration tool with multi-tenant schema support.

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
