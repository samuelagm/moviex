FROM golang:alpine AS builder

# Minimal build dependencies
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy only the dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build with CGO disabled (static binary)
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /app

# Create non-root user
COPY --from=builder /etc/passwd /etc/passwd
USER nobody

EXPOSE 8080

CMD ["/app"]
