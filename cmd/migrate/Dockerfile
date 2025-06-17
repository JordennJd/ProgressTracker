# Используем образ Go
FROM golang:latest as builder

WORKDIR /app

# Копируем зависимости и собираем мигратор
COPY go.mod go.sum ./
RUN go mod download

COPY "cmd/migrate/main.go" .
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate ./main.go

# Финальный образ (минимальный)
FROM alpine:latest

RUN apk --no-cache add postgresql-client

WORKDIR /app
COPY --from=builder /migrate /app/migrate
COPY migrations /app/migrations

ENTRYPOINT ["/app/migrate"]