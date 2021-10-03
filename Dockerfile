# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY . ./
RUN go mod vendor
RUN go build -o fizzbuzz ./cmd/fizzbuzz/main.go 

EXPOSE 8888

CMD ["/app/fizzbuzz", "-p", "8888"]