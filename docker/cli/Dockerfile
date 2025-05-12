# Etapa de build com imagem Alpine
FROM golang:1.24.1-alpine AS builder

# Instala git e outras dependências (necessário para baixar módulos)
RUN apk add --no-cache git

WORKDIR /app

# Copia os arquivos e resolve as dependências
COPY . .
COPY .env ./
COPY go.mod go.sum ./
RUN go mod download

# Compila o binário com saída para /app/cli
RUN go build -o cli ./cmd/cli

ENTRYPOINT ["./cli"]