# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first (agar cache dependency bisa lebih optimal)
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .
# COPY .env .env

# Build the application
RUN go build -o app

# Final stage
FROM alpine:latest

# Install SSL certificates (wajib buat HTTPS request)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from builder
COPY --from=builder /app/app .
# COPY --from=builder /app/.env .env


# Command to run
CMD ["./app"]




