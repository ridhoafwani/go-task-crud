FROM golang:1.23-alpine as builder

# Set working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .
COPY .env .

# Build the Go app
RUN go build -o main cmd/main.go

# Use a smaller image for the final container
FROM alpine:latest

RUN apk --no-cache add bash

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
