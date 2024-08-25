# Send Email Consumer

This application is a Kafka consumer written in Go that listens to a specific Kafka topic for incoming messages containing email information. Upon receiving a message, the application sends an email using the Resend email service.

## Table of Contents

- [Project Structure](#project-structure)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [Building the Docker Image](#building-the-docker-image)
- [Deploying the Application](#deploying-the-application)

## Project Structure

- `main.go`: The main application file that contains the logic to consume Kafka messages and send emails.
- `Dockerfile`: A multi-stage Dockerfile to build and run the application.
- `deploy.yml`: GitHub Actions workflow file to automate deployment to a remote server.
- `go.mod`, `go.sum`: Go modules files that manage dependencies.

## Features

- **Kafka Consumer**: Listens to a Kafka topic for incoming messages.
- **Email Sending**: Sends an email using the Resend email service when a message is received.
- **Dockerized**: The application is containerized using Docker for easy deployment.
- **Automated Deployment**: Uses GitHub Actions to automate deployment to a remote server.

## Requirements

- Go 1.20 or later
- Docker
- Kafka broker
- Resend email service API key
- Access to a Linux-based server for deployment

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/send-email-consumer.git
    cd send-email-consumer
    ```

2. Set up your environment variables as described below.

## Environment Variables

The application requires several environment variables to run. These can be set in your deployment environment (e.g., in the Docker container) or passed directly during runtime.

- `RESEND_API_KEY`: Your API key for the Resend email service.
- `KAFKA_BROKER`: The address of your Kafka broker.
- `KAFKA_TOPIC`: The Kafka topic to which the consumer will subscribe.
- `KAFKA_GROUP_ID`: The consumer group ID for Kafka.
- `KAFKA_SASL_USERNAME`: SASL username for Kafka authentication.
- `KAFKA_SASL_PASSWORD`: SASL password for Kafka authentication.

## Building the Docker Image

The application is designed to be built and run in a Docker container. You can build the Docker image using the following command:

```bash
docker buildx build --platform linux/arm64 -t send-email-consumer:latest --load .
