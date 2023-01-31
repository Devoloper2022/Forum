////////////////Hgs

FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -o main ./cmd/main.go
FROM ubuntu:20.04
WORKDIR /app

COPY --from=builder /app .
EXPOSE 4000
CMD ["./main"]