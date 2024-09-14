# Send Email Consumer

This application is a RabbitMQ consumer written in Go that listens to a specific RabbitMQ queue for incoming messages containing email information. Upon receiving a message, the application sends an email using the Resend email service.

## Table of Contents

- [Project Structure](#project-structure)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [Building the Docker Image](#building-the-docker-image)
- [Deploying the Application](#deploying-the-application)

## Project Structure

- `main.go`: The main application file that contains the logic to consume RabbitMQ messages and send emails.
- `Dockerfile`: A multi-stage Dockerfile to build and run the application.
- `deploy.yml`: GitHub Actions workflow file to automate deployment to a remote server.
- `go.mod`, `go.sum`: Go modules files that manage dependencies.

## Features

- **RabbitMQ Consumer**: Listens to a RabbitMQ queue for incoming messages.
- **Email Sending**: Sends an email using the Resend email service when a message is received.
- **Dockerized**: The application is containerized using Docker for easy deployment.
- **Automated Deployment**: Uses GitHub Actions to automate deployment to a remote server.

## Requirements

- Go 1.20 or later
- Docker
- RabbitMQ broker
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
- `RABBITMQ_URL`: The connection URL for RabbitMQ in the format `amqp://user:password@host:port/vhost`.
- `RABBITMQ_QUEUE`: The RabbitMQ queue to which the consumer will subscribe.
  
### Optional Environment Variables
- `RABBITMQ_EXCHANGE`: The RabbitMQ exchange name (if needed for advanced routing).
- `RABBITMQ_ROUTING_KEY`: The routing key used if subscribing to a topic exchange in RabbitMQ.

## Building the Docker Image

The application is designed to be built and run in a Docker container. You can build the Docker image using the following command:

```bash
docker buildx build --platform linux/amd64 -t send-email-consumer:latest --load .
