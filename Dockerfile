# syntax=docker/dockerfile:1

FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./
RUN mkdir ./fetcher
ADD fetcher ./fetcher
RUN mkdir ./models
ADD models ./models

RUN go build -o /fetch

ENTRYPOINT [ "/fetch" ]
