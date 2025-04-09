FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main

ENTRYPOINT ["/app/main"]
