# Multi-stage build for Kalco
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
ARG VERSION=dev
ARG COMMIT=unknown
ARG DATE=unknown

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w -X main.version=$VERSION -X main.commit=$COMMIT -X main.date=$DATE" \
    -o kalco .

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata git

# Create non-root user
RUN addgroup -g 1001 -S kalco && \
    adduser -u 1001 -S kalco -G kalco

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/kalco /app/kalco

# Change ownership to non-root user
RUN chown -R kalco:kalco /app

# Switch to non-root user
USER kalco

# Expose port (if needed for future web interface)
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["/app/kalco"]

# Default command
CMD ["--help"]
