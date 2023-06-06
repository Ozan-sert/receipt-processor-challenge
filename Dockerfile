# Start from a clean, minimal base image with Go installed
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download dependencies (improves caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose the desired port(s) for the application
EXPOSE 8080

# Set the command to run the Go application
CMD ["./main"]
