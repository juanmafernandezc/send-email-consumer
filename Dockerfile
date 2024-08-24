FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /send-email-consumer main.go

FROM gcr.io/distroless/base-debian10

COPY --from=builder /send-email-consumer /send-email-consumer

CMD ["/send-email-consumer"]