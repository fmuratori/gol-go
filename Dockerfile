# Use the official golang image as the base image
FROM golang:1.22.2-alpine3.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN chmod +x build.sh 
RUN ./build.sh

# Start a new stage from scratch
FROM alpine:3.19

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/golgo .

# Command to run the executable
CMD ["./golgo"]
