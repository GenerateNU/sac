# syntax=docker/dockerfile:1

FROM golang:1.22.2

# Copy files and download dependencies
RUN mkdir /app
ADD . /app
WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

# Build and make executable
RUN CGO_ENABLED=0 GOOS=linux go build -o /sac
RUN chmod +x /sac

# Install migrate for migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run command
WORKDIR /
CMD ["/sac"]