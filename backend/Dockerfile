# Builder Stage
FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Development/Runtime Stage
FROM golang:1.24.4-alpine

RUN apk add --no-cache ca-certificates libc6-compat

# air for live reload
RUN go install github.com/air-verse/air@latest

WORKDIR /app

EXPOSE 3000
