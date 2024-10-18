FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main

FROM alpine:latest

 WORKDIR /app

 COPY --from=builder /app/main .

 EXPOSE 9901

 CMD ["./main"]
 