package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/resend/resend-go/v2"
	"github.com/streadway/amqp"
)

type EmailMessage struct {
	ReplyTo string `json:"replyTo"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func main() {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	queueName := os.Getenv("RABBITMQ_QUEUE")
	resendAPIKey := os.Getenv("RESEND_API_KEY")

	if rabbitURL == "" || queueName == "" || resendAPIKey == "" {
		log.Fatal("Faltan variables de entorno necesarias")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("No se pudo conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("No se pudo abrir un canal: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("No se pudo declarar la cola: %v", err)
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("No se pudo registrar un consumidor: %v", err)
	}

	client := resend.NewClient(resendAPIKey)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recibido mensaje: %s", d.Body)

			err := sendEmail(client, d.Body)
			if err != nil {
				log.Printf("Error enviando email: %v", err)
			}
		}
	}()

	log.Printf("Esperando mensajes en la cola %s. Para salir presiona CTRL+C", queueName)
	<-forever
}

func sendEmail(client *resend.Client, message []byte) error {
	var emailMessage EmailMessage

	err := json.Unmarshal(message, &emailMessage)
	if err != nil {
		return fmt.Errorf("error deserializando mensaje: %w", err)
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
		return fmt.Errorf("error enviando email: %w", err)
	}

	log.Printf("Email enviado, ID: %s\n", sent.Id)
	return nil
}
