FROM golang:1.22.6-alpine3.20 AS builder
RUN apk update
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go build -o kraitlog-api ./cmd/api

FROM alpine:3.20 as binary
COPY --from=builder /app/kraitlog-api .
EXPOSE 8080
CMD ["./kraitlog-api"]