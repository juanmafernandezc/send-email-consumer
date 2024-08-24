FROM golang:1.20-alpine as builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /send-email-consumer main.go

FROM alpine:3.18

RUN apk add --no-cache libc6-compat

COPY --from=builder /send-email-consumer /send-email-consumer

CMD ["/send-email-consumer"]