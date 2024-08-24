# Dockerfile usando Buildx para crear imágenes multi-arch
FROM --platform=$BUILDPLATFORM golang:1.20-alpine AS builder

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

# Instalar dependencias necesarias para compilación estática
RUN apk add --no-cache gcc musl-dev

# Copiar el archivo go.mod y go.sum
COPY go.mod .
COPY go.sum .

# Descargar las dependencias
RUN go mod download

# Copiar el resto del código de la aplicación
COPY . .

# Compilar la aplicación de manera estática
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /send-email-consumer -ldflags="-s -w" main.go

# Usar una imagen minimalista que soporte ARM64 para la ejecución
FROM --platform=$TARGETPLATFORM arm64v8/alpine:3.18

# Copiar el binario compilado desde la etapa de construcción
COPY --from=builder /send-email-consumer /send-email-consumer

# Definir el comando por defecto para ejecutar la aplicación
CMD ["/send-email-consumer"]
