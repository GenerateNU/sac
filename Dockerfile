# syntax=docker/dockerfile:1

FROM golang:1.22.2

RUN mkdir /app
ADD ./backend /app
WORKDIR /app

COPY ./backend/go.mod ./backend/go.sum ./

# Download all the dependencies
RUN go mod download

COPY ./backend .
COPY ./config /config

RUN CGO_ENABLED=0 GOOS=linux go build -o /sac
RUN chmod +x /sac

WORKDIR /
CMD ["/sac", "-config=/config"]