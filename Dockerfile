#
# Stage 1: Build the Go application
#
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o ./hellok3sglaucus .
#
# Stage 2: Minimal image for running the app
#
FROM alpine:latest

COPY --from=builder /app/hellok3sglaucus /usr/bin/hellok3sglaucus
EXPOSE 80

CMD ["hellok3sglaucus"]
