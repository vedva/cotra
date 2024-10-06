# Step 1: Build the Go app in a builder container
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests and install dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY ./cmd/main.go ./cmd/

# Build the Go app and name the executable 'cotra'
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /cotra ./cmd/main.go

# Step 2: Copy the compiled app into a smaller image
FROM alpine:latest

# Set up a working directory
WORKDIR /root/

# Copy the binary from the builder container to the final container
COPY --from=builder /cotra .

# Expose port 8080 for the service
EXPOSE 8080

# Command to run the binary
CMD ["./cotra"]
