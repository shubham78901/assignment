# Build stage
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .  
RUN go mod download && go build -o main ./api/cmd/main.go

# Final minimal image
FROM alpine:latest
COPY --from=builder /app/main /main

EXPOSE 8000
CMD ["/main"]
