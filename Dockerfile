# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS builder

WORKDIR /app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

COPY . .


RUN go build -o main main.go



FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

COPY .env .
EXPOSE 8080
CMD [ "/app/main" ]