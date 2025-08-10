# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum init.sql ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Stage 2: Create the final image
# FROM alpine:latest
# WORKDIR /app
# COPY --from=builder ./cmd/ .
EXPOSE 8080
CMD ["./main"]