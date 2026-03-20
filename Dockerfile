# STAGE 1: The Builder
FROM golang:1.26-alpine AS builder

# Install git/ca-certificates (Alpine needs these for private modules/HTTPS)
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Cache dependencies first (The "Speed" trick)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd

# STAGE 2: The Final Runtime
FROM scratch

# Import certificates from builder (Scratch has no root CA)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the static binary
COPY --from=builder /app/server /server

# Run as non-root for security
USER 1001

ENTRYPOINT ["/server"]