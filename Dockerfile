FROM arm64v8/debian:stretch-slim as builder

WORKDIR /app

RUN apt-get update && apt-get install -y build-essential

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /send-email-consumer -ldflags="-s -w" main.go

FROM arm64v8/debian:stretch-slim

COPY --from=builder /app/send-email-consumer /send-email-consumer

CMD ["/send-email-consumer"]
