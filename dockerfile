# Use the official Go image with the specified version
FROM golang:1.23.2 AS builder

# Set environment variables
ENV GO111MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o app .

# Use a minimal image for running the application
FROM alpine:latest

# Install SQLite and CA certificates
RUN apk --no-cache add sqlite-libs ca-certificates

# Set working directory in the container
WORKDIR /root/

# Copy the built binary and necessary files
COPY --from=builder /app/app .
COPY db/database.db ./db/database.db

# Expose the port the app runs on
EXPOSE 8080

# Run the app
CMD ["./app"]
