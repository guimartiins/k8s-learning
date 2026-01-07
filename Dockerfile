FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY ./server.go .

RUN go build -o server server.go

FROM alpine:latest

COPY --from=builder /app/server /usr/local/bin/server

CMD ["server"]