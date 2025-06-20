FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Compile go app into a single binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o /beathub-backend .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /beathub-backend .

EXPOSE 8080

CMD ["./beathub-backend"]

