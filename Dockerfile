FROM debian:buster-slim as builder

RUN apt-get update && apt-get install -y \
    golang-go \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /send-email-consumer main.go

FROM debian:buster-slim

RUN apt-get update && apt-get install -y \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /send-email-consumer /send-email-consumer

CMD ["/send-email-consumer"]
