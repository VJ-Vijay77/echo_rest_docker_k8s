# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o binary main.go



FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/binary /app/
COPY .env .
EXPOSE 8080
CMD [ "/app/binary" ]