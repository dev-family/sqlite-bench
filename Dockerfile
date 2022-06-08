FROM golang:1.18-alpine AS builder

RUN apk add build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

#RUN #CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go


