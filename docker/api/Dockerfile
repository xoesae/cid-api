FROM golang:1.24.1-alpine AS base

FROM base AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 8080

CMD ["air"]

FROM base AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o cid-api cmd/api/main.go

# Production
FROM scratch AS production

WORKDIR /prod

COPY --from=builder /build/cid-api ./
COPY --from=builder /build/.env ./

EXPOSE 8080

CMD ["/prod/cid-api"]
