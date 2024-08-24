package main

import (
    "context"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/segmentio/kafka-go"
    "github.com/segmentio/kafka-go/sasl/scram"
    "github.com/resend/resend-go/v2"
)

type EmailMessage struct {
    ReplyTo string `json:"replyTo"`
    To      string `json:"to"`
    Subject string `json:"subject"`
    Message string `json:"message"`
}

func Consumer(topic string, groupId string, resendAPIKey string) {
    saslUsername := os.Getenv("KAFKA_SASL_USERNAME")
    saslPassword := os.Getenv("KAFKA_SASL_PASSWORD")
    kafkaBroker := os.Getenv("KAFKA_BROKER")

    mechanism, err := scram.Mechanism(scram.SHA512, saslUsername, saslPassword)
    if err != nil {
        log.Fatalf("Failed to create SASL mechanism: %v", err)
    }

    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{kafkaBroker},
        GroupID: groupId,
        Topic:   topic,
        Dialer: &kafka.Dialer{
            SASLMechanism: mechanism,
            TLS:           &tls.Config{},
        },
    })

    defer r.Close()

    client := resend.NewClient(resendAPIKey)

    for {
        ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
        defer cancel()

        message, err := r.ReadMessage(ctx)
        if err != nil {
            log.Printf("Error reading message: %v", err)
            continue
        }

        fmt.Printf("Message received: Partition: %d Offset: %d Value: %s\n", message.Partition, message.Offset, string(message.Value))

        err = sendEmail(client, message.Value)
        if err != nil {
            log.Printf("Error sending email: %v", err)
        }
    }
}

func sendEmail(client *resend.Client, message []byte) error {
    var emailMessage EmailMessage

    err := json.Unmarshal(message, &emailMessage)
    if err != nil {
        return fmt.Errorf("error deserializing message: %w", err)
    }

    params := &resend.SendEmailRequest{
        From:    "Juanmalabs <develop@juanmalabs.com>",
        To:      []string{emailMessage.To},
        Html:    emailMessage.Message,
        Subject: emailMessage.Subject,
        ReplyTo: emailMessage.ReplyTo,
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        return fmt.Errorf("error sending email: %w", err)
    }

    log.Printf("Email sent, ID: %s\n", sent.Id)
    return nil
}

func main() {
    resendAPIKey := os.Getenv("RESEND_API_KEY")
    topic := os.Getenv("KAFKA_TOPIC")
    groupId := os.Getenv("KAFKA_GROUP_ID")

    if resendAPIKey == "" || topic == "" || groupId == "" {
        log.Fatal("Missing required environment variables")
    }

    Consumer(topic, groupId, resendAPIKey)
}