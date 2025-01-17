# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Install Git (required for some dependencies)
RUN apk add --no-cache git

ENV GOPROXY=https://proxy.golang.org,direct

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/server

# Deployment stage
FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add libc6-compat
COPY --from=builder /app/main .
COPY --from=builder /app/internal/config/pack_sizes.json ./internal/config/pack_sizes.json
EXPOSE 8080
CMD ["./main"]