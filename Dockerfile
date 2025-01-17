# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /backend

# Install Git (required for some dependencies)
RUN apk add --no-cache git

# Copy go.mod and go.sum first to cache dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the entire source code
COPY backend/ .

# Build the Go application
RUN go build -o main ./cmd/server

# Deployment stage
FROM alpine:latest
WORKDIR /
RUN apk --no-cache add libc6-compat
COPY --from=builder /backend/main .
COPY --from=builder /backend/internal/config/pack_sizes.json ./internal/config/pack_sizes.json
EXPOSE 8080
CMD ["./main"]
