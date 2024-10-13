FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go


FROM alpine:3.18


WORKDIR /app

COPY --from=builder /app/main .


COPY .env .env
COPY config ./config


EXPOSE 4000


CMD ["./main"]
