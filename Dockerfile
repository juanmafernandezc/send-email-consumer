FROM arm64v8/golang:1.20 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOARCH=arm64 go build -o /send-email-consumer main.go

FROM arm64v8/ubuntu:20.04

RUN apt-get update && apt-get install -y \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /send-email-consumer /send-email-consumer

CMD ["/send-email-consumer"]
