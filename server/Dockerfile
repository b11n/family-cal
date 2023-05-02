# syntax=docker/dockerfile:1

FROM golang:1.20.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

ADD / ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /web-server

EXPOSE 8080

CMD ["/web-server"]