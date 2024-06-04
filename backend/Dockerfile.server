# syntax=docker/dockerfile:1

FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./sac

EXPOSE 8080

CMD ["/sac"]