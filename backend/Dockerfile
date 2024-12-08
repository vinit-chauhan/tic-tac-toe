# Use the official Golang image as the build environment
FROM golang:1.23-bookworm AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.* .
RUN go mod download

# Copy the rest of the application code
COPY config ./config
COPY . .

# Set environment variables for the build
ENV CGO_ENABLED=0 
ENV GOOS=linux

# Build the Go application
RUN go build -o build/server ./cmd/server

# Use a minimal Alpine image for the final container
FROM alpine:3.18

WORKDIR /app

# Copy the configuration file
COPY config.yml .

# Copy the built server binary from the builder stage
COPY --from=builder /app/build/server .
