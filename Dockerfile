FROM golang:1.20-alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o send-email-consumer main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/send-email-consumer .

CMD ["./send-email-consumer"]
