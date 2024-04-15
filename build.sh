#!/bin/sh

# Initialize the Go module and generate the go.mod file
go mod init golgo

# Download dependencies (if any)
go mod tidy

# Build the Go binary
go build -o golgo main.go
