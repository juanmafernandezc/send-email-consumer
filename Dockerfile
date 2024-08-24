FROM golang:1.20-alpine as builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /send-email-consumer -ldflags="-s -w" main.go

FROM arm64v8/alpine:3.18

COPY --from=builder /send-email-consumer /send-email-consumer

CMD ["/send-email-consumer"]
