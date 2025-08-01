# Etapa 1: Compilar la app
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o login-service cmd/main.go

# Etapa 2: Imagen ligera
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/login-service .

EXPOSE 8181
CMD ["./login-service"]
