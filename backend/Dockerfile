# syntax=docker/dockerfile:1

# Use a smaller base image for production
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go binary (optimized for production)
RUN mkdir -p /app/bin
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/bin/sac

# Install migrate for migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

USER nobody

# Start a new stage to create the final production image
FROM alpine:3.19.1
WORKDIR /app

# Copy only the necessary binary from the builder stage
COPY --from=builder /app/bin/sac /app/sac

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl -f http://localhost:8080/health || exit 1

CMD ["/app/sac"]
