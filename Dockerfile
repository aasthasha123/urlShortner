# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .

# 🔥 IMPORTANT: build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

CMD ["./app"]