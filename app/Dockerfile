FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/info cmd/info/main.go

FROM alpine:latest AS runtime

WORKDIR /app
COPY --from=builder /app/bin/info ./