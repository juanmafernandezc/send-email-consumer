FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN GOOS=linux GOARCH=arm64 go build -o /send-email-consumer -ldflags="-s -w" main.go

FROM gcr.io/distroless/base-debian11

COPY --from=builder /send-email-consumer /send-email-consumer

CMD ["/send-email-consumer"]
