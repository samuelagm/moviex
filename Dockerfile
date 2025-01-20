# Start with the official Golang Alpine image
FROM golang:alpine AS builder

# Install necessary packages
RUN apk update && apk add --no-cache git gcc libc-dev

# Set the working directory for the build
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o /go/bin/app

# Start a new minimal image
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk add --no-cache libc6-compat

# Copy the binary from the builder stage
COPY --from=builder /go/bin/app /app

# Expose port 8080 (adjust if necessary for your application)
EXPOSE 8080

# Command to run the application
CMD ["/app"]
