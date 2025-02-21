
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY api ./api

WORKDIR /app/api
RUN go build -o main ./cmd/main.go

# Final minimal image
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/api/main .

EXPOSE 8000

CMD ["./main"]