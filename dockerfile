# Use the official Go image as a builder
FROM golang:1.23.2 as builder

# Set the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal base image for production
FROM alpine:latest

# Install SQLite
RUN apk --no-cache add sqlite

# Set the working directory
WORKDIR /app

# Copy the SQLite database
COPY --from=builder /app/job.db /app/job.db

# Copy the built application
COPY --from=builder /app/main /app/main

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
