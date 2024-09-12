FROM golang:1.22.4-alpine3.19

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -g cmd/api/main.go -o docs

RUN go mod tidy